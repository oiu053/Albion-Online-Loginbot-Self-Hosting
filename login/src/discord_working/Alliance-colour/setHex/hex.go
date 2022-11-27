package sethex

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

func Alliance_colour_set_hex(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	//Content1 := cmd.ApplicationCommandData().Options[0].Value

	var ContentR int
	var ContentG int
	var ContentB int
	var err error

	option := cmd.ApplicationCommandData().Options[0].Value
	Hex := fmt.Sprintf("value: %v", option)

	Hex = strings.TrimPrefix(Hex, "value: ")

	c, err := operation.ParseHexColor(Hex)

	if err != nil {
		return
	}

	ContentR = int(c.R)
	ContentG = int(c.G)
	ContentB = int(c.B)

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
