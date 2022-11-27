package unregistercmd

import (
	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

func Unregister(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	//Content := strings.TrimPrefix(cmd.Content, cmdreg)

	//fmt.Println("t1")

	Check := xmldb.Check_for_file(cmd.GuildID, cmd.Member.User.ID)

	//fmt.Print(Check)
	if !Check {
		//fmt.Println("Error Nr. 1")
		s.ChannelMessageSend(cmd.ChannelID, "You are not registered!")
		return
	}

	e1, errors := xmldb.Xmlfile_read(cmd.GuildID, cmd.Member.User.ID)

	if errors {
		s.ChannelMessageSend(cmd.ChannelID, "Something went wrong. Please try again!")
		return
	}

	Check = xmldb.Check_for_file_txt(cmd.GuildID, e1.Playername)
	//fmt.Print(Check)
	if !Check {
		s.ChannelMessageSend(cmd.ChannelID, "This acc is not registered!")
		return
	}

	OK := xmldb.DeleatefilesofthePlayer(cmd.GuildID, e1.Playername)

	if !OK {
		s.ChannelMessageSend(cmd.ChannelID, "Something went wrong. Please try again!")
		return
	}

	RoleIDS := cmd.Member.Roles
	for _, RoleID := range RoleIDS {
		s.GuildMemberRoleRemove(cmd.GuildID, cmd.Member.User.ID, RoleID)
	}
	s.ChannelMessageSend(cmd.ChannelID, "You are unregistered")
	s.GuildMemberNickname(cmd.GuildID, cmd.Member.User.ID, "")
	return
}
