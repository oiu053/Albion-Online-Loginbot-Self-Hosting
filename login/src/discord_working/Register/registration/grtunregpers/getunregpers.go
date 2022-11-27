package grtunregpers

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	albionAPI "github.com/oiu053/Login_Bot_Albion/src/Albion_Api"
)

var storragefolder string = "storrage/"

func Get_unregisteredplayers(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

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

	//loading(s, cmd)

	content1 := cmd.ApplicationCommandData().Options[0].Value

	content := fmt.Sprintf("%v", content1)

	var Players []fs.DirEntry
	for _, v := range direntery {
		if strings.HasSuffix(v.Name(), ".txt") {
			Players = append(Players, v)
		}
	}

	//s.ChannelMessageSend(cmd.ChannelID, strconv.Itoa(len(Players))+" registered player!")

	Playernames := []string{}

	for _, v := range Players {
		Playernames = append(Playernames, strings.TrimSuffix(v.Name(), ".txt"))
	}

	//fmt.Print(albionAPI.SearchGuildID(content))
	Membernames := albionAPI.GetguildMembers_by_ID(albionAPI.SearchGuildID(content))

	//fmt.Print(content)

	Unregisterednames := []string{}

	for _, v := range Membernames {
		Newmembername := true
		for _, v2 := range Playernames {
			if Newmembername {
				if strings.ToLower(v) == strings.ToLower(v2) {
					Newmembername = false
				}
			}
		}
		if Newmembername {
			Unregisterednames = append(Unregisterednames, v)
		}
	}

	//fmt.Print(Unregisterednames)

	Unregisterednamesstring := ""

	for _, v := range Unregisterednames {
		Unregisterednamesstring = Unregisterednamesstring + "\n" + v
	}

	z, _ := s.ChannelMessageSend(cmd.ChannelID, Unregisterednamesstring)

	s.ChannelMessagePin(cmd.ChannelID, z.ID)

	s.ChannelMessageSend(cmd.ChannelID, "The List willnot get deleated automaticly cause of the Pinned status!")

}
