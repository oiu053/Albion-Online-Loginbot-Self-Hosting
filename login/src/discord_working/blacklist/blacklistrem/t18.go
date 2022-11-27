package blacklistrem

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

var storragefolder string = "storrage/"

func Blacklist_remove(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	p, err := s.UserChannelPermissions(cmd.Member.User.ID, cmd.ChannelID)
	if err != nil {
		return
	}

	if discordgo.PermissionBanMembers != p&discordgo.PermissionBanMembers {
		s.ChannelMessageSend(cmd.ChannelID, "You do not have the required permission!")
		return
	}

	Content1 := cmd.ApplicationCommandData().Options[0].Value

	//fmt.Println(Content1)

	Content := fmt.Sprintf("value: %v", Content1)
	//fmt.Println("t1")
	Content = strings.TrimPrefix(Content, "value: ")

	exists := xmldb.Check_for_file_blacklist(cmd.GuildID)

	if !exists {
		s.ChannelMessageSend(cmd.ChannelID, "Noone of your server is blacklisted!")
		return
	} else {

		e1, errors := xmldb.ReadBlacklist(cmd.GuildID)

		if errors {
			return
		}

		var Blacklisted []string

		for _, v := range e1.Blacklist {
			if strings.ToLower(v) != strings.ToLower(Content) {
				Blacklisted = append(Blacklisted, v)
			}
		}

		err1 := xmldb.WriteBlacklist(cmd.GuildID, Blacklisted, e1.Blacklistedguilds)

		if !err1 {
			s.ChannelMessageSend(cmd.ChannelID, "Something went Wrong")
			return
		}
		s.ChannelMessageSend(cmd.ChannelID, "The player got blacklist_removeed!!")

		//MesRef := cmd.Message.Reference()

		//_, err := s.ChannelMessageSendReply(cmd.ChannelID, "Test!", MesRef)

		//fmt.Print(err)

	}
}
