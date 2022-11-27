package operations

import (
	"context"
	"errors"
	"fmt"
	"image/color"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	"golang.org/x/sync/singleflight"
)

// CreateRole is a RequestType enumeration.
const CreateRole RequestType = iota

// RequestType string representations.
const (
	CreateRoleString = "CreateRole"
	UnknownString    = "unknown"
)

// APIErrorCodeMaxRoles is the Discord API error code for max roles.
const APIErrorCodeMaxRoles = 91703936005

const (
	roleHoist             = true
	roleMention           = true
	guildMembersPageLimit = 1000
)

// Request is an operations request to be processed.
type Request struct {
	Type       RequestType
	CreateRole *CreateRoleRequest
}

// RequestType represents a type of operations request.
type RequestType int

// String returns the string representation for the given RequestType.
func (rt RequestType) String() string {
	switch rt {
	case CreateRole:
		return CreateRoleString
	default:
		return UnknownString
	}
}

// CreateRoleRequest is a request to create a new role.
type CreateRoleRequest struct {
	Guild     *discordgo.Guild
	RoleName  string
	RoleColor int
}

// Gateway is a centralized construct to process operation requests by
// de-duplicating identical simultaneous requests and providing the result to
// all of the callers.
type Gateway struct {
	Session *discordgo.Session
	group   *singleflight.Group
}

// NewGateway returns a new *Gateway ready to process requests.
func NewGateway(session *discordgo.Session) *Gateway {
	return &Gateway{
		Session: session,
		group:   &singleflight.Group{},
	}
}

// Process will process the provided request and send back the result to the
// provided ResultChannel. The caller should type check the result it receives
// to determine if an error was sent or the result is of the type it expects.
func (gateway *Gateway) Process(request *Request) <-chan singleflight.Result {
	key := fmt.Sprintf("%s/%s/%s",
		request.Type,
		request.CreateRole.Guild.ID,
		request.CreateRole.RoleName,
	)

	defer gateway.group.Forget(key)

	return gateway.group.DoChan(key, func() (any, error) {
		switch request.Type {
		case CreateRole:
			return createRole(
				gateway.Session,
				request.CreateRole.Guild,
				request.CreateRole.RoleName,
				request.CreateRole.RoleColor,
			)
		default:
			return nil, fmt.Errorf("%s request type not supported", request.Type)
		}
	})
}

// LookupGuild returns a *discordgo.Guild from the session's internal state
// cache. If the guild is not found in the state cache, LookupGuild will query
// the Discord API for the guild and add it to the state cache before returning
// it.
func LookupGuild(session *discordgo.Session, guildID string) (*discordgo.Guild, error) {
	guild, err := session.State.Guild(guildID)
	if err != nil {
		guild, err = updateStateGuilds(session, guildID)
		if err != nil {
			return nil, fmt.Errorf("unable to query guild: %w", err)
		}
	}

	return guild, nil
}

// AddRoleToMember adds the role associated with the provided roleID to the
// user associated with the provided userID, in the guild associated with the
// provided guildID.
func AddRoleToMember(session *discordgo.Session, guildID, userID, roleID string) error {
	err := session.GuildMemberRoleAdd(guildID, userID, roleID)
	if err != nil {
		return fmt.Errorf("unable to add ephemeral role: %w", err)
	}

	return nil
}

// RemoveRoleFromMember removes the role associated with the provided roleID
// from the user associated with the provided userID, in the guild associated
// with the provided guildID.
func RemoveRoleFromMember(session *discordgo.Session, guildID, userID, roleID string) error {
	err := session.GuildMemberRoleRemove(guildID, userID, roleID)
	if err != nil {
		return fmt.Errorf("unable to remove ephemeral role: %w", err)
	}

	return nil
}

// IsDeadlineExceeded checks if the provided error wraps
// context.DeadlineExceeded.
func IsDeadlineExceeded(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}

// IsForbiddenResponse checks if the provided error wraps *discordgo.RESTError.
// If it does, IsForbiddenResponse returns true if the response code is equal
// to http.StatusForbidden.
func IsForbiddenResponse(err error) bool {
	var restErr *discordgo.RESTError

	if errors.As(err, &restErr) {
		if restErr.Response.StatusCode == http.StatusForbidden {
			return true
		}
	}

	return false
}

// IsMaxGuildsResponse checks if the provided error wraps *discordgo.RESTError.
// If it does, IsMaxGuildsResponse returns true if the response code is equal
// to http.StatusBadRequest and the error code is 91703936005.
func IsMaxGuildsResponse(err error) bool {
	var restErr *discordgo.RESTError

	if errors.As(err, &restErr) {
		if restErr.Response.StatusCode == http.StatusBadRequest {
			return restErr.Message.Code == APIErrorCodeMaxRoles
		}
	}

	return false
}

// ShouldLogDebug checks if the provided error should be logged at a debug
// level.
func ShouldLogDebug(err error) bool {
	switch {
	case IsDeadlineExceeded(err), IsForbiddenResponse(err):
		return true
	default:
		return false
	}
}

// BotHasChannelPermission checks if the bot has view permissions for the
// channel. If the bot does have the view permission, BotHasChannelPermission
// returns nil.
func BotHasChannelPermission(session *discordgo.Session, channel *discordgo.Channel) error {
	permissions, err := session.UserChannelPermissions(session.State.User.ID, channel.ID)
	if err != nil {
		return fmt.Errorf("unable to determine channel permissions: %w", err)
	}

	if permissions&discordgo.PermissionViewChannel != discordgo.PermissionViewChannel {
		return fmt.Errorf("insufficient channel permissions: channel: %s", channel.Name)
	}

	return nil
}

func updateStateGuilds(session *discordgo.Session, guildID string) (*discordgo.Guild, error) {
	guild, err := session.Guild(guildID)
	if err != nil {
		return nil, fmt.Errorf("error sending guild query request: %w", err)
	}

	roles, err := session.GuildRoles(guildID)
	if err != nil {
		return nil, fmt.Errorf("unable to query guild channels: %w", err)
	}

	channels, err := session.GuildChannels(guildID)
	if err != nil {
		return nil, fmt.Errorf("unable to query guild channels: %w", err)
	}

	members, err := recursiveGuildMembers(session, guildID, "", guildMembersPageLimit)
	if err != nil {
		return nil, fmt.Errorf("unable to query guild members: %w", err)
	}

	guild.Roles = roles
	guild.Channels = channels
	guild.Members = members
	guild.MemberCount = len(members)

	err = session.State.GuildAdd(guild)
	if err != nil {
		return nil, fmt.Errorf("unable to add guild to state cache: %w", err)
	}

	return guild, nil
}

func createRole(
	session *discordgo.Session,
	guild *discordgo.Guild,
	roleName string,
	roleColor int,
) (*discordgo.Role, error) {
	role, err := session.GuildRoleCreate(guild.ID, &discordgo.RoleParams{
		Name:        roleName,
		Color:       &roleColor,
		Hoist:       pointerTo(roleHoist),
		Mentionable: pointerTo(roleMention),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create ephemeral role: %w", err)
	}

	err = session.State.RoleAdd(guild.ID, role)
	if err != nil {
		return nil, fmt.Errorf("unable to add ephemeral role to state cache: %w", err)
	}

	return role, nil
}

func Createrole(
	session *discordgo.Session,
	guild *discordgo.Guild,
	roleName string,
	roleColor int,
) (*discordgo.Role, error) {
	role, err := session.GuildRoleCreate(guild.ID, &discordgo.RoleParams{
		Name:        roleName,
		Color:       &roleColor,
		Hoist:       pointerTo(roleHoist),
		Mentionable: pointerTo(roleMention),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create ephemeral role: %w", err)
	}

	err = session.State.RoleAdd(guild.ID, role)
	if err != nil {
		return nil, fmt.Errorf("unable to add ephemeral role to state cache: %w", err)
	}

	return role, nil
}

func recursiveGuildMembers(
	session *discordgo.Session,
	guildID, after string,
	limit int,
) ([]*discordgo.Member, error) {
	guildMembers, err := session.GuildMembers(guildID, after, limit)
	if err != nil {
		return nil, fmt.Errorf("error sending recursive guild members request: %w", err)
	}

	if len(guildMembers) < guildMembersPageLimit {
		return guildMembers, nil
	}

	nextGuildMembers, err := recursiveGuildMembers(
		session,
		guildID,
		guildMembers[len(guildMembers)-1].User.ID,
		guildMembersPageLimit,
	)
	if err != nil {
		return nil, err
	}

	return append(guildMembers, nextGuildMembers...), nil
}

func pointerTo[T any](v T) *T {
	return &v
}
func PointerTo[T any](v T) *T {
	return &v
}

func Roleexists(Rolename string, s *discordgo.Session, g *discordgo.Guild) (Role *discordgo.Role, Exists bool) {
	for _, Role := range g.Roles {
		Checkname := Role.Name
		if strings.ToLower(Rolename) == strings.ToLower(Checkname) {
			return Role, true
		}
	}

	return &discordgo.Role{}, false
}
func Delateroleby_playername(s *discordgo.Session, GuildID, Playername string) {

	exists := xmldb.Check_for_file_txt(GuildID, Playername)

	if exists {

		DiscordID, errors := xmldb.GetDiscordPlayerID_by_IGN(GuildID, Playername)
		if errors {
			fmt.Print("Error z1")
			return
		}

		Guilds := s.State.Guilds

		for _, v := range Guilds {
			if v.ID == GuildID {
				for _, member := range v.Members {
					if member.User.ID == DiscordID {
						for _, Role := range member.Roles {
							s.GuildMemberRoleRemove(GuildID, DiscordID, Role)

						}
						return
					}
				}
			}
		}

	}
}
func Delateroleby_ID(s *discordgo.Session, GuildID, IDs string) {

	Guilds := s.State.Guilds

	for _, v := range Guilds {
		if v.ID == GuildID {
			for _, member := range v.Members {
				if member.User.ID == IDs {
					for _, Role := range member.Roles {
						s.GuildMemberRoleRemove(GuildID, IDs, Role)
					}
					return
				}
			}
		}
	}

}

func RoleColourIntCreator(ContentR, ContentG, ContentB int) (Colour int) {
	Rechnung1 := ContentR * 65536
	Rechnung2 := ContentG * 256
	Rechnung3 := ContentB

	Colour = Rechnung1 + Rechnung2 + Rechnung3

	return Colour
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func Commandanswer(s *discordgo.Session, cmd *discordgo.InteractionCreate, Command string) {
	Importantconst := Command

	s.InteractionRespond(cmd.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Importantconst,
		},
	})

}

func Folder_exists(URL string) (exists bool) {
	_, err := os.Stat(URL)
	if err != nil {
		if err == nil {
			return true
		}
		return false
	}
	return true
}

func Create_Folder(URL string) (created bool) {
	error := os.MkdirAll(URL, os.ModePerm)
	if error != nil {
		error = os.MkdirAll(URL, os.ModePerm)
		if error != nil {
			return false
		}
	}
	return true
}