package registerchanneladd

import (
	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

var storragefolder string = "storrage/"

func Registerchannelset(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	e1, errors := xmldb.Xmlfileconfig_read(cmd.GuildID)

	if errors {
		return
	}

	//permissions /execute ajhsgdajhvfahsdkagsdkjagskhdbakhsfkjha
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

	GU, err := s.Guild(cmd.GuildID)
	if err != nil {
		s.ChannelMessageSend(cmd.ChannelID, "An Error had occured. Please try again!")
		return
	}

	succsess := xmldb.Xmlfileconfig_create(cmd.GuildID, GU, e1.AdminRoles, e1.AdminIGNs, e1.GuildMemberRolesIDs, e1.FriendsIngamename, e1.FriendRolesIDs, cmd.ChannelID, e1.RoleColour, e1.AllianceRoleColour)

	if succsess {
		s.ChannelMessageSend(cmd.ChannelID, "The Registerchannel is set to <#"+cmd.ChannelID+">")
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "An Error had occured. Please try again!")
}
