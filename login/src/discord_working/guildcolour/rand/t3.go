package rand

import (
	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

var storragefolder string = "storrage/"

func Guildcolourrand(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
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

	var Nilrolecolour int = 917039360
	e1.RoleColour = Nilrolecolour

	succses := xmldb.Xmlfileconfig_create(GuildID, Guild, e1.AdminRoles, e1.AdminIGNs, e1.GuildMemberRolesIDs, e1.FriendsIngamename, e1.FriendRolesIDs, e1.RegisterchannelID, e1.RoleColour, e1.AllianceRoleColour)
	if succses {
		s.ChannelMessageSend(cmd.ChannelID, "Config Had been edited!")
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "Please try again")
	return
}
