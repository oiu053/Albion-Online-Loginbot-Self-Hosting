package timestampcreator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Timestamp(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	a := cmd.ApplicationCommandData().Options

	var year int
	var month int
	var day int
	var hour int
	var minute int

	for _, v := range a {
		if v.Name == "yyyy" {
			year = integer_to_int(v.Value)
		}
		if v.Name == "mm" {
			month = integer_to_int(v.Value)
		}
		if v.Name == "dd" {
			day = integer_to_int(v.Value)
		}
		if v.Name == "hh" {
			hour = integer_to_int(v.Value)
		}
		if v.Name == "minmin" {
			minute = integer_to_int(v.Value)
		}
	}

	/*abziehen1 := 1969 * 31536000
	abziehen2 := 11 * 2628000
	abziehen3 := 24 * 86400 * 0
	abziehen := abziehen1 + abziehen2 + abziehen3
	year = year * 31536000
	month = month * 2628000
	day = day * 86400
	hour = hour * 3600
	minute = minute * 60
	*/
	date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	timestamp := fmt.Sprintf("%v", date.Unix())
	timestampint, err := strconv.Atoi(timestamp)
	if err != nil {
		return
	}

	timestampint = timestampint - 3600

	timestamp = strconv.Itoa(timestampint)

	//strconv.Itoa(timestamp)

	s.ChannelMessageSend(cmd.ChannelID, "<t:"+timestamp+":R>"+"\n<t:"+timestamp+">"+` / \<t:`+timestamp+`:R>`+` / \<t:`+timestamp+`>`)
}

func integer_to_int(a interface{}) (integer int) {

	v := fmt.Sprintf("%v", a)

	int, err := strconv.Atoi(v)

	if err == nil {
		return int
	}

	return 0
}
