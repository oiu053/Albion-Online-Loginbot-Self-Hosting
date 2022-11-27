package checkfwrongregistrations

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

func Check_f_wrong_registrations(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	IDs, errors := check_file_player(s, cmd.GuildID)

	if errors {
		s.ChannelMessageSend(cmd.ChannelID, "An error had occured")

		stringtosend := "false Registerations: "
		for _, v := range IDs {
			stringtosend = stringtosend + " <@" + v + ">"
		}
		z, err := s.ChannelMessageSend(cmd.ChannelID, stringtosend)
		if err != nil {
			return
		}
		s.ChannelMessagePin(cmd.ChannelID, z.ID)
		return
	}
	stringtosend := "false Registerations: "
	for _, v := range IDs {
		stringtosend = stringtosend + " <@" + v + ">"
	}
	z, err := s.ChannelMessageSend(cmd.ChannelID, stringtosend)
	if err != nil {
		return
	}
	s.ChannelMessagePin(cmd.ChannelID, z.ID)
	return
}

func check_file_player(s *discordgo.Session, GuildID string) (DiscordIDs []string, successfull bool) {
	filename := storragefolder + GuildID
	readed, err := os.ReadDir(filename)

	if err != nil {
		return []string{}, false
	}

	checkable_files := []fs.DirEntry{}
	for _, file := range readed {
		if strings.HasSuffix(file.Name(), ".xml") {
			checkable_files = append(checkable_files, file)
		}
	}

	for _, v := range checkable_files {
		Playerinfo, errors := xmldb.Xmlfile_read(GuildID, strings.TrimSuffix(v.Name(), ".xml"))

		if errors {
			return DiscordIDs, false
		}

		alliancename, errors2 := albionAPI.SearchPlayerAllianceName(Playerinfo.Playername)

		if errors2 {
			return DiscordIDs, false
		}
		guildname := albionAPI.SearchGuild(Playerinfo.Playername)

		if strings.ToLower(guildname) != strings.ToLower(Playerinfo.Guildname) {
			ID, errors := Forceunregister(s, GuildID, Playerinfo)
			if errors {
				return DiscordIDs, false
			}
			DiscordIDs = append(DiscordIDs, ID)
		} else if strings.ToLower(alliancename) != strings.ToLower(Playerinfo.AllianceName) {
			ID, errors := Forceunregister(s, GuildID, Playerinfo)
			if errors {
				return DiscordIDs, false
			}
			DiscordIDs = append(DiscordIDs, ID)
		}
	}
	return DiscordIDs, true
}

func Forceunregister(s *discordgo.Session, GuildID string, Playerinfo xmldb.Xml_struct_Player_info) (ID string, errors bool) {
	operation.Delateroleby_playername(s, GuildID, Playerinfo.Playername)
	xmldb.DeleatefilesofthePlayer(GuildID, Playerinfo.Playername)
	s.GuildMemberNickname(GuildID, Playerinfo.DiscordUserId, "Forceunregistered by: system")
	return Playerinfo.DiscordUserId, false
}
