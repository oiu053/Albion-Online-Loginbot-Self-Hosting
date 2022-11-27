package getregisteredpersons

import (
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var storragefolder string = "storrage/"

func Get_registered_players(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

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

	var Players []fs.DirEntry
	for _, v := range direntery {
		if strings.HasSuffix(v.Name(), ".txt") {
			Players = append(Players, v)
		}
	}

	s.ChannelMessageSend(cmd.ChannelID, strconv.Itoa(len(Players))+" registered player!")

	Playernames := []string{}

	for _, v := range Players {
		Playernames = append(Playernames, strings.TrimSuffix(v.Name(), ".txt"))
	}

	Playernamestringlist := ""

	for _, v := range Playernames {
		Playernamestringlist = Playernamestringlist + "\n" + v
	}

	z, err := s.ChannelMessageSend(cmd.ChannelID, Playernamestringlist)

	if err != nil {
		return
	}
	err = s.ChannelMessagePin(cmd.ChannelID, z.ID)

	if err != nil {
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "The List willnot get deleated automaticly cause of the Pinned status!")

}
