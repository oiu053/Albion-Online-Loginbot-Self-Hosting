package set

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

func Alliance_colour_set(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	//Content1 := cmd.ApplicationCommandData().Options[0].Value

	var ContentR int
	var ContentG int
	var ContentB int
	var err error

	for _, option := range cmd.ApplicationCommandData().Options {
		if option.Name == "r" {

			ContentR1 := option.Value
			ContentR2 := fmt.Sprintf("value: %v", ContentR1)

			ContentR, err = strconv.Atoi(strings.TrimPrefix(ContentR2, "value: "))

			if err != nil {
				return
			}
			//ColourR =
		} else if option.Name == "g" {

			ContentG1 := option.Value
			ContentG2 := fmt.Sprintf("value: %v", ContentG1)

			ContentG, err = strconv.Atoi(strings.TrimPrefix(ContentG2, "value: "))

			if err != nil {
				return
			}
			//ColourG =
		} else if option.Name == "b" {

			ContentB1 := option.Value
			ContentB2 := fmt.Sprintf("value: %v", ContentB1)

			ContentB, err = strconv.Atoi(strings.TrimPrefix(ContentB2, "value: "))

			if err != nil {
				return
			}
			//ColourG =
		}

	}

	GuildID := cmd.GuildID

	e1, errors := xmldb.Xmlfileconfig_read(GuildID)
	if errors {

		return
	}

	Execute := false
	for _, Role := range cmd.Member.Roles {
		for _, v := range e1.AdminRoles {
			if Role == v {
				Execute = true
			}
		}
	}

	p, err := s.UserChannelPermissions(cmd.Member.User.ID, cmd.ChannelID)
	if err != nil {
		return
	}

	if discordgo.PermissionManageServer == p&discordgo.PermissionManageServer {
		Execute = true
	}
	if env.BotOwner() == cmd.Member.User.ID {
		Execute = true
	}

	if !Execute {
		s.ChannelMessageSend(cmd.ChannelID, "You need to have Adminpermissions!")
		return
	}

	if errors {
		s.ChannelMessageSend(cmd.ChannelID, "Please try again")
		return
	}

	Guild, err := s.Guild(GuildID)
	if err != nil {
		s.ChannelMessageSend(cmd.ChannelID, "Please try again")
		return
	}

	e1.AllianceRoleColour = operation.RoleColourIntCreator(ContentR, ContentG, ContentB)

	succses := xmldb.Xmlfileconfig_create(GuildID, Guild, e1.AdminRoles, e1.AdminIGNs, e1.GuildMemberRolesIDs, e1.FriendsIngamename, e1.FriendRolesIDs, e1.RegisterchannelID, e1.RoleColour, e1.AllianceRoleColour)
	if succses {
		s.ChannelMessageSend(cmd.ChannelID, "Config Had been edited!")
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "Please try again")
	return
}
