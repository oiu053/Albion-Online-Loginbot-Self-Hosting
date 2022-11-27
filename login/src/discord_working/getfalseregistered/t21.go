package getfalseregistered

import (
	"io/fs"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	albionAPI "github.com/oiu053/Login_Bot_Albion/src/Albion_Api"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

var storragefolder string = "storrage/"

func Get_false_registered(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	////loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	//permissions
	Execute := false

	Guld, err005 := s.Guild(cmd.GuildID)
	if err005 != nil {
		return
	}

	if cmd.Member.User.ID == Guld.OwnerID {
		Execute = true
	}

	if !Execute {
		s.ChannelMessageSend(cmd.ChannelID, "You need to be the serverowner to use this command!")
		return
	}
	direntery, error003 := os.ReadDir(storragefolder + cmd.GuildID)
	if error003 != nil {
		return
	}

	////loading(s, cmd)

	var Players []fs.DirEntry
	for _, v := range direntery {
		if strings.HasSuffix(v.Name(), ".xml") {
			Players = append(Players, v)
		}
	}

	//s.ChannelMessageSend(cmd.ChannelID, strconv.Itoa(len(Players))+" registered player!")

	PlayerIDS := []string{}

	for _, v := range Players {
		PlayerIDS = append(PlayerIDS, strings.TrimSuffix(v.Name(), ".xml"))
	}

	_, PlayerIDS_to_UNREGISTER, e4 := check_if_correctregistered(PlayerIDS, cmd.GuildID) //Playernames1

	if e4 {
		s.ChannelMessageSend(cmd.ChannelID, "An error had occuered")
		return
	}

	var PINGMESSAGE string
	for _, v := range PlayerIDS_to_UNREGISTER {
		PINGMESSAGE = PINGMESSAGE + "<@" + v + ">\n"
	}

	s.ChannelMessageSendComplex(cmd.ChannelID, &discordgo.MessageSend{
		Content: PINGMESSAGE,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "force-unregister",
						Style: discordgo.DangerButton,
						Emoji: discordgo.ComponentEmoji{
							Name: "🚷",
						},
						CustomID: "forceunregister",
					},
				},
			},
		},
	})

	//if err != nil {
	/*
		for _, v := range Playernames1 {
			forceunregister(s, cmd, v)
		}
	*/
	//}
	//s.ChannelMessagePin(cmd.ChannelID, stn.ID)

}

func check_if_correctregistered(PlayerIDS []string, GuildID string) (PlayerNames_to_unregister, PlayerIDS_to_unregister []string, errors bool) {

	for _, PlID := range PlayerIDS {
		data, e1 := xmldb.Xmlfile_read(GuildID, PlID)

		if e1 {
			return []string{}, []string{}, true
		}

		var Newplayer bool = true

		PlaydataGuildname, e2 := albionAPI.SearchPlayerGuildName(data.Playername)
		PlaydataAlliname, e3 := albionAPI.SearchPlayerAllianceName(data.Playername)

		if e2 {
			return []string{}, []string{}, true
		}
		if e3 {
			return []string{}, []string{}, true
		}

		if Newplayer {
			if strings.ToLower(data.Guildname) != strings.ToLower(PlaydataGuildname) {
				PlayerNames_to_unregister = append(PlayerIDS_to_unregister, data.Playername)
				PlayerIDS_to_unregister = append(PlayerIDS_to_unregister, data.DiscordUserId)
				Newplayer = false
			}
		}
		if Newplayer {
			if strings.ToLower(data.AllianceName) != strings.ToLower(PlaydataAlliname) {
				PlayerNames_to_unregister = append(PlayerIDS_to_unregister, data.Playername)
				PlayerIDS_to_unregister = append(PlayerIDS_to_unregister, data.DiscordUserId)
				Newplayer = false
			}
		}
	}
	return PlayerNames_to_unregister, PlayerIDS_to_unregister, false
}

func forceunregister(s *discordgo.Session, cmd *discordgo.InteractionCreate, PlayerName string) {

	////loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	Content := PlayerName

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

		PlayerID, errors1 := xmldb.GetDiscordPlayerID_by_IGN(cmd.GuildID, Content)

		if errors1 {
			msg, err := s.ChannelMessageSend(cmd.ChannelID, "An Error had Occured: "+PlayerName+" <@"+PlayerID+">")
			if err == nil {
				s.ChannelMessagePin(cmd.ChannelID, msg.ID)
			}
			return
		}

		/*
			if env.BotOwner() == PlayerID{
				Execute = false
			}
		*/

		operation.Delateroleby_playername(s, cmd.GuildID, PlayerName)

		msg, err := s.ChannelMessageSend(cmd.ChannelID, "The player <@"+PlayerID+"> now unregistered!"+PlayerName)
		xmldb.DeleatefilesofthePlayer(cmd.GuildID, Content)
		s.GuildMemberNickname(cmd.GuildID, PlayerID, "Forceunregistered by: "+cmd.Member.Nick+"/"+cmd.Member.User.Username)
		if err == nil {
			s.ChannelMessagePin(cmd.ChannelID, msg.ID)
		}
		return

		//

	}

}
