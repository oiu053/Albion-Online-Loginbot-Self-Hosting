package funcforceunregisterbuttonanswer

import (
	"github.com/bwmarrin/discordgo"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
)

func Forceunregister(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	//Permission:
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
	edited := discordgo.MessageEdit{
		Content: &cmd.Message.Content,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "force-unregister",
						Style: discordgo.DangerButton,
						Emoji: discordgo.ComponentEmoji{
							Name: "❌",
						},
						Disabled: true,
						CustomID: "forceunregister",
					},
				},
			},
		},
		ID:      cmd.Message.ID,
		Channel: cmd.ChannelID,
	}

	s.ChannelMessageEditComplex(&edited)
	s.ChannelMessageUnpin(cmd.ChannelID, cmd.ID)

	Users := cmd.Message.Mentions

	var Playernamesrowping string
	for _, v := range Users {
		e1, error1 := xmldb.Xmlfile_read(cmd.GuildID, v.ID)

		if error1 {
			s.ChannelMessageSend(cmd.ChannelID, "Error")
			return
		}

		PlayerID := forceunregister(s, cmd, e1.Playername)

		Playernamesrowping = Playernamesrowping + "\n<@" + PlayerID + ">"
	}

	operation.Commandanswer(s, cmd, "_ _")
	msg, err3 := s.ChannelMessageSend(cmd.ChannelID, "The following players have been forceunregistered: "+Playernamesrowping)

	if err3 == nil {
		s.ChannelMessagePin(cmd.ChannelID, msg.ID)
	}
}

func forceunregister(s *discordgo.Session, cmd *discordgo.InteractionCreate, PlayerName string) (PlayerID_if_true string) {

	////loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return ""
	}

	Content := PlayerName

	//fmt.Print("t7")

	/*Blacklisted := []string{Content}

	err1 := xmldb.WriteBlacklist(cmd.GuildID, Blacklisted)

	if !err1 {
		s.ChannelMessageSend(cmd.ChannelID,  "Something went Wrong")
		return
	}
	//operation.Delateroleby_playername(s, cmd.GuildID, Content)
	s.ChannelMessageSend(cmd.ChannelID,  "The player is blacklisted!")*/

	PlayerID, errors1 := xmldb.GetDiscordPlayerID_by_IGN(cmd.GuildID, Content)

	if errors1 {
		s.ChannelMessageSend(cmd.ChannelID, "An Error had Occured: "+PlayerName+" <@"+PlayerID+">")
		//if err == nil {
		//	s.ChannelMessagePin(cmd.ChannelID, msg.ID)
		//}
		return ""
	}

	/*
		if env.BotOwner() == PlayerID{
			Execute = false
		}
	*/

	operation.Delateroleby_playername(s, cmd.GuildID, PlayerName)

	//s.ChannelMessageSend(cmd.ChannelID, "The player <@"+PlayerID+"> now unregistered!"+PlayerName)
	xmldb.DeleatefilesofthePlayer(cmd.GuildID, Content)
	s.GuildMemberNickname(cmd.GuildID, PlayerID, "Forceunregistered by: "+cmd.Member.Nick+"/"+cmd.Member.User.Username)
	//if err == nil {
	//	s.ChannelMessagePin(cmd.ChannelID, msg.ID)
	//}

	//

	return PlayerID

}
