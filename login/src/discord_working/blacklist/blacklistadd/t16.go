package blacklistadd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

var storragefolder string = "storrage/"

func Blacklist(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	Content1 := cmd.ApplicationCommandData().Options[0].Value

	//fmt.Println(Content1)

	Content := fmt.Sprintf("value: %v", Content1)
	//fmt.Println("t1")
	Content = strings.TrimPrefix(Content, "value: ")

	//fmt.Print("t7")
	exists := xmldb.Check_for_file_blacklist(cmd.GuildID)

	if !exists {
		/*Blacklisted := []string{Content}

		err1 := xmldb.WriteBlacklist(cmd.GuildID, Blacklisted)

		if !err1 {
			s.ChannelMessageSend(cmd.ChannelID,  "Something went Wrong")
			return
		}
		//operation.Delateroleby_playername(s, cmd.GuildID, Content)
		s.ChannelMessageSend(cmd.ChannelID,  "The player is blacklisted!")*/
	} else {

		e1, errors := xmldb.ReadBlacklist(cmd.GuildID)

		if errors {
			return
		}
		config, errors := xmldb.Xmlfileconfig_read(cmd.GuildID)

		if errors {
			return
		}

		Execute := true
		for _, Role := range cmd.Member.Roles {
			for _, v := range config.AdminRoles {
				if Role == v {
					Execute = false
				}
			}
		}

		p, err := s.UserChannelPermissions(cmd.Member.User.ID, cmd.ChannelID)
		if err != nil {
			return
		}

		if discordgo.PermissionManageServer == p&discordgo.PermissionManageServer {
			Execute = false
		}

		GU, err2 := s.Guild(cmd.GuildID)

		if err2 != nil {
			s.ChannelMessageSend(cmd.ChannelID, "Please try Again!")
			return
		}

		if GU.OwnerID == cmd.Member.User.ID {
			Execute = true
		}

		PlayerID, errors1 := xmldb.GetDiscordPlayerID_by_IGN(cmd.GuildID, Content)

		if errors1 {
			return
		}

		//fmt.Print(PlayerID)
		if GU.OwnerID == PlayerID {
			Execute = false
		}
		/*
			if env.BotOwner() == PlayerID{
				Execute = false
			}
		*/

		if env.BotOwner() == cmd.Member.User.ID {
			Execute = true
		}

		if !Execute {
			s.ChannelMessageSend(cmd.ChannelID, "The character who should get blacklisted has admin permissions!")
			return
		}

		if discordgo.PermissionBanMembers != p&discordgo.PermissionBanMembers {
			s.ChannelMessageSend(cmd.ChannelID, "You do not have the required permission!")
			return
		}

		operation.Delateroleby_playername(s, cmd.GuildID, Content)

		Blacklisted := append(e1.Blacklist, Content)

		err1 := xmldb.WriteBlacklist(cmd.GuildID, Blacklisted, e1.Blacklistedguilds)

		if !err1 {
			s.ChannelMessageSend(cmd.ChannelID, "Something went Wrong")
			return
		} else {
			s.ChannelMessageSend(cmd.ChannelID, "This player is now blacklisted!")
			xmldb.DeleatefilesofthePlayer(cmd.GuildID, Content)
			s.GuildMemberNickname(cmd.GuildID, PlayerID, "Blacklisted by: "+cmd.Member.Nick+"/"+cmd.Member.User.Username)
			return
		}
		//

	}

}
