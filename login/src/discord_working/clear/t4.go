package clear

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

//var storragefolder string = "storrage/"

func Clear1(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//Permission

	e1, errors3 := xmldb.Xmlfileconfig_read(cmd.GuildID)

	if errors3 {
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
		//loading(s, cmd)
		s.ChannelMessageSend(cmd.ChannelID, "You need to have Adminpermissions!")
		return
	}

	Messages, err := s.ChannelMessages(cmd.ChannelID, 100, "", "", "")
	//loading(s, cmd)

	if err != nil {
		return
	}

	var Deleated_Messages int = 0

	//fmt.Print(Messages)

	for _, m := range Messages {

		if !m.Pinned {
			err2 := s.ChannelMessageDelete(cmd.ChannelID, m.ID)

			if err2 == nil {
				Deleated_Messages = Deleated_Messages + 1
			}

		}

	}

	s.ChannelMessageSend(cmd.ChannelID, strconv.Itoa(Deleated_Messages)+" from "+strconv.Itoa(len(Messages))+" Messages got deleated!")

}

func Clear_all2(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//Permission

	e1, errors3 := xmldb.Xmlfileconfig_read(cmd.GuildID)

	if errors3 {
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
		//loading(s, cmd)
		s.ChannelMessageSend(cmd.ChannelID, "You need to have Adminpermissions!")
		return
	}

	Messages, err := s.ChannelMessages(cmd.ChannelID, 100, "", "", "")
	//loading(s, cmd)
	if err != nil {
		return
	}

	var MessagesIDs []string
	for _, v := range Messages {
		MessagesIDs = append(MessagesIDs, v.ID)
	}

	s.ChannelMessagesBulkDelete(cmd.ChannelID, MessagesIDs)

	s.ChannelMessageSend(cmd.ChannelID, "Messages got deleated!")

}
