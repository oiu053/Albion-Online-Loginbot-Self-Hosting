package adminroleadd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

var storragefolder string = "storrage/"

func Adminroles_add(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	Content1 := cmd.ApplicationCommandData().Options[0].Value

	//fmt.Println(Content1)

	Content := fmt.Sprintf("value: %v", Content1)
	//fmt.Println("t1")
	Content = strings.TrimPrefix(Content, "value: ")

	//fmt.Println(Content)
	GuildID := cmd.GuildID

	e1, errors := xmldb.Xmlfileconfig_read(GuildID)
	if errors {

		return
	}

	Execute := false
	/*for _, Role := range cmd.Member.Roles {
		for _, v := range e1.AdminRoles {
			if Role == v {
				Execute = true
			}
		}
	}*/
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
	RoleID := strings.TrimSuffix(strings.TrimPrefix(Content, "<@&"), ">")

	e1.AdminRoles = append(e1.AdminRoles, RoleID)

	succses := xmldb.Xmlfileconfig_create(GuildID, Guild, e1.AdminRoles, e1.AdminIGNs, e1.GuildMemberRolesIDs, e1.FriendsIngamename, e1.FriendRolesIDs, e1.RegisterchannelID, e1.RoleColour, e1.AllianceRoleColour)
	if succses {
		s.ChannelMessageSend(cmd.ChannelID, "Config Had been edited!")
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "Please try again")
	return
}
