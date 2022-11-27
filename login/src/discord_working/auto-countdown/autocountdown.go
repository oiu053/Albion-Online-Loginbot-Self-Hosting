package autocountdown

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Autocontent(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	a := cmd.ApplicationCommandData().Options

	//timevar := time.Now().Unix()

	var day int
	var hour int
	var minute int

	for _, v := range a {
		New := true
		if New {
			if v.Name == "day" {
				day = integer_to_int(v.Value)
			}
		}
		if New {
			if v.Name == "hour" {
				hour = integer_to_int(v.Value)
			}
		}
		if New {
			if v.Name == "minute" {
				minute = integer_to_int(v.Value)
			}
		}
	}

	dayadd := day * 86400
	if day != 0 {
		dayadd = dayadd + 3600
	}
	//dayadd := dayadd1 + 3600 + 3600
	houradd := hour * 3600
	minuteadd := minute * 60

	timef := int(time.Now().UTC().Unix())
	timestamp := timef + dayadd + houradd + minuteadd
	//strconv.Itoa(timestamp)

	s.ChannelMessageSend(cmd.ChannelID, "<t:"+strconv.Itoa(timestamp)+":R>"+` / \<t:`+strconv.Itoa(timestamp)+":R>")
}

func integer_to_int(a interface{}) (integer int) {

	v := fmt.Sprintf("%v", a)

	int, err := strconv.Atoi(v)

	if err == nil {
		return int
	}

	return 0
}
