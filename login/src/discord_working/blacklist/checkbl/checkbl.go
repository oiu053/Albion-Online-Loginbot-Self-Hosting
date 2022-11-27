package checkbl

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

var storragefolder string = "storrage/"

func Checkbl(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)
	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	Content1 := cmd.ApplicationCommandData().Options[0].Value

	//fmt.Println(Content1)

	Content := fmt.Sprintf("value: %v", Content1)
	//fmt.Println("t1")
	Content = strings.TrimPrefix(Content, "value: ")

	a, errors := xmldb.ReadBlacklist(cmd.GuildID)

	if errors {
		s.ChannelMessageSend(cmd.ChannelID, "An Error with the blacklist check had occured. please try again later!")
		return
	}

	Blacklist := a.Blacklist

	for _, Blacklisted_Pl := range Blacklist {
		if strings.ToLower(Blacklisted_Pl) == strings.ToLower(Content) {
			s.ChannelMessageSend(cmd.ChannelID, "The character: "+Content+" is Blacklisted")
			return
		}
	}
	s.ChannelMessageSend(cmd.ChannelID, "The character: "+Content+" is **not** Blacklisted")
	return
}
