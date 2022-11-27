package xmldb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Blacklistfilename string = "/Blacklistfile.xml"
var storragefolder string = "storrage/"
var trashfolder string = "!trash/"
var Logfolder string = "log/"

//var Filename string = "storrage/" + "GuildID_PlayerID.xml"

type Xml_struct_Player_info struct {
	Playername               string `xml:"playername"`
	Guildname                string `xml:"guildname"`
	AllianceName             string `xml:"alliancename"`
	DiscordUserId            string `xml:"discorduserid"`
	DiscordUserEmail         string `xml:"discorduseremail"`
	DiscordUserEmailverified bool   `xml:"discorduseremailverified"`
	DiscordPlayersLanguage   string `xml:"discorduserlanguage"`
	DiscordMFAEnabled        bool   `xml:"discordmfaenabled"`
	DiscordHash              string `xml:"discorduserhash"`
	Blacklisted              bool   `xml:"blacklist"`
}
type xml_struct_config struct {
	ServerName    string `xml:"servername"`
	ServerID      string `xml:"serverid"`
	ServerOwnerID string `xml:"serverownerid"`

	AdminRoles []string `xml:"adminroles"`
	AdminIGNs  []string `xml:"adminigns"`

	FriendsIngamename []string `xml:"friendsign"`
	FriendRolesIDs    []string `xml:"friendrole"`

	GuildMemberRolesIDs []string `xml:"memberrole"`

	RegisterchannelID string `xml:"registerchannelid"`
	///Implement!!!!!!
	RoleColour         int `xml:"rolecolour"`
	AllianceRoleColour int `xml:"alliancerolecolour"`
}
type xml_blacklist_struct struct {
	Blacklist         []string `xml:"blacklistedplayer"`
	Blacklistedguilds []string `xml:"blacklistedguilds"`
}

func filename(GuildID string, PlayerID string) (filename string) {
	return storragefolder + GuildID + "/" + PlayerID + ".xml"
}
func filenameLog(GuildID, ChannelID string) (filename string) {
	return Logfolder + GuildID + "/" + ChannelID + ".txt"
}

func filenameConfig(GuildID string) (filename string) {
	return storragefolder + GuildID + ".xml"
}

func filenamePlayer(GuildID string, PlayerName string) (filename string) {
	return storragefolder + GuildID + "/" + strings.ToLower(PlayerName) + ".txt"
}

func Check_for_file(GuildID string, PlayerID string) (exists bool) {

	filename := filename(GuildID, PlayerID)

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}
func Check_for_file_Guildlog_txt(GuildID, ChannelID string) (exists bool) {

	filename := filenameLog(GuildID, ChannelID)

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}
func Check_for_file_txt(GuildID string, Playername string) (exists bool) {

	filename := filenamePlayer(GuildID, Playername)

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}
func Check_for_file_blacklist(GuildID string) (exists bool) {

	filename := "./blacklisted/" + GuildID + Blacklistfilename

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}

func check_for_file_filename(filename string) (exists bool) {

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}

func exists_createfolder(GuildID string) (OK bool) {

	foldername := storragefolder + GuildID

	if _, err := os.Stat(foldername); err == nil {
		return true
	} else {
		err := os.MkdirAll(foldername, 0755)
		if err != nil {
			return false
		}

		return true
	}

}
func exists_createfolder_trash(GuildID string) (OK bool) {

	foldername := trashfolder + GuildID

	if _, err := os.Stat(foldername); err == nil {
		return true
	} else {
		err := os.MkdirAll(foldername, 0755)
		if err != nil {
			return false
		}

		return true
	}

}
func exists_createfolder_for_BL(GuildID string) (OK bool) {

	foldername := "./blacklisted/" + GuildID

	if _, err := os.Stat(foldername); err == nil {
		return true
	} else {
		err := os.MkdirAll(foldername, 0755)
		if err != nil {
			return false
		}

		return true
	}

}
func exists_createfoldername(foldername string) (OK bool) {

	if _, err := os.Stat(foldername); err == nil {
		return true
	} else {
		err := os.MkdirAll(foldername, 0755)
		if err != nil {
			return false
		}

		return true
	}

}

func Xmlfile_create_manual(GuildID string, DiscordsUserID string, PlayersIGN string, PlayersIGGuildName string, DiscordEmail string, Discordemailverify bool, DiscordsPlayerslanguage string, DiscordMFA bool, DiscordHash string, Blacklist bool, AllianceName string) (succses bool) {

	Check1 := exists_createfolder(GuildID)

	if !Check1 {
		return false
	}
	filename := filename(GuildID, DiscordsUserID)

	e1 := &Xml_struct_Player_info{
		Playername:               PlayersIGN,
		Guildname:                PlayersIGGuildName,
		AllianceName:             AllianceName,
		DiscordUserId:            DiscordsUserID,
		DiscordUserEmail:         DiscordEmail,
		DiscordUserEmailverified: Discordemailverify,
		DiscordPlayersLanguage:   DiscordsPlayerslanguage,
		DiscordMFAEnabled:        DiscordMFA,
		DiscordHash:              DiscordHash,
		Blacklisted:              Blacklist,
	}

	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return false
	}

	err1 := create_Name_to_Player(GuildID, PlayersIGN, DiscordsUserID)

	if err1 {
		return false
	}

	return true
}

func Xmlfile_create_auto(GuildID string, PlayerID string, Playername string, Guildname string, AllianceName string, m *discordgo.MessageCreate) (succses bool) {

	Check1 := exists_createfolder(GuildID)
	if !Check1 {
		return false
	}

	filename := filename(GuildID, PlayerID)

	e1 := &Xml_struct_Player_info{
		Playername:               Playername,
		Guildname:                Guildname,
		AllianceName:             AllianceName,
		DiscordUserId:            m.Author.ID,
		DiscordUserEmail:         m.Author.Email,
		DiscordUserEmailverified: m.Author.Verified,
		DiscordPlayersLanguage:   m.Author.Locale,
		DiscordMFAEnabled:        m.Author.MFAEnabled,
		DiscordHash:              m.Author.Discriminator,
		Blacklisted:              false,
	}

	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return false
	}

	err1 := create_Name_to_Player(GuildID, Playername, m.Author.ID)

	if err1 {
		return false
	}

	return true
}
func Xmlfile_create_auto_cmd(GuildID string, PlayerID string, Playername string, Guildname string, AllianceName string, cmd *discordgo.InteractionCreate) (succses bool) {

	Check1 := exists_createfolder(GuildID)
	if !Check1 {
		return false
	}

	filename := filename(GuildID, PlayerID)

	e1 := &Xml_struct_Player_info{
		Playername:               Playername,
		Guildname:                Guildname,
		AllianceName:             AllianceName,
		DiscordUserId:            cmd.Member.User.ID,
		DiscordUserEmail:         cmd.Member.User.Email,
		DiscordUserEmailverified: cmd.Member.User.Verified,
		DiscordPlayersLanguage:   cmd.Member.User.Locale,
		DiscordMFAEnabled:        cmd.Member.User.MFAEnabled,
		DiscordHash:              cmd.Member.User.Discriminator,
		Blacklisted:              false,
	}

	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return false
	}

	err1 := create_Name_to_Player(GuildID, Playername, cmd.Member.User.ID)

	if err1 {
		return false
	}

	return true
}

func Xmlfile_read(GuildID string, PlayerID string) (file Xml_struct_Player_info, errors bool) {

	filename := filename(GuildID, PlayerID)

	e1 := &Xml_struct_Player_info{}
	Check := check_for_file_filename(filename)

	if !Check {
		return *e1, true
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return *e1, true
	}

	err = xml.Unmarshal([]byte(data), &e1)
	if err != nil {
		return *e1, true
	}

	return *e1, false
}

func create_Name_to_Player(GuildID string, IGN string, PlayerID string) (errors bool) {

	filename := filenamePlayer(GuildID, IGN)
	f, err := os.Create(filename)

	if err != nil {
		return true
	}

	defer f.Close()

	_, err2 := f.WriteString(PlayerID)

	if err2 != nil {
		return true
	}

	return false
}

func GetDiscordPlayerID_by_IGN(GuildID string, PlayerName string) (DiscordPlayerID string, errors bool) {

	content, err := os.ReadFile(filenamePlayer(GuildID, PlayerName))
	if err != nil {
		return "", true
	}
	PlayerID := string(content)
	return PlayerID, false
}

func DeleatefilesoftheGuild(GuildID string) (OK bool) {

	now := time.Now()

	newfolder := strconv.Itoa(now.UTC().Day()) + strconv.Itoa(now.UTC().Hour()) + strconv.Itoa(now.UTC().Minute()) + strconv.Itoa(now.UTC().Second()) + strconv.Itoa(now.UTC().Year()) + now.UTC().Weekday().String() + "_" + GuildID
	/*ok := exists_createfolder_trash(newfolder)

	if !ok {
		return false
	}*/

	err := os.Rename("./storrage/"+GuildID, "./!trash/"+newfolder+"info")
	if err != nil {
		//fmt.Println(err)
	}
	err = os.Rename("./blacklisted/"+GuildID, "./!trash/"+strconv.Itoa(now.UTC().Day())+strconv.Itoa(now.UTC().Hour())+strconv.Itoa(now.UTC().Minute())+strconv.Itoa(now.UTC().Second())+strconv.Itoa(now.UTC().Year())+now.UTC().Weekday().String()+"_bl_"+GuildID+".xml")
	if err != nil {
		//fmt.Println(err)

	}
	err = os.Rename("./storrage/"+GuildID+".xml", "./!trash/"+strconv.Itoa(now.UTC().Day())+strconv.Itoa(now.UTC().Hour())+strconv.Itoa(now.UTC().Minute())+strconv.Itoa(now.UTC().Second())+strconv.Itoa(now.UTC().Year())+now.UTC().Weekday().String()+GuildID+"_old.xml")
	if err != nil {
		fmt.Print(err)
	}

	return true
}
func DeleatefilesofthePlayer(GuildID string, IGN string) (OK bool) {
	a, check := GetDiscordPlayerID_by_IGN(GuildID, IGN)

	if check {
		return false
	}

	now := time.Now()

	ok := exists_createfolder_trash(GuildID)

	if !ok {
		return false
	}

	err := os.Rename("./storrage/"+GuildID+"/"+a+".xml", "./!trash/"+GuildID+"/"+strconv.Itoa(now.UTC().Day())+strconv.Itoa(now.UTC().Hour())+strconv.Itoa(now.UTC().Minute())+strconv.Itoa(now.UTC().Second())+strconv.Itoa(now.UTC().Year())+now.UTC().Weekday().String()+"_old_"+a+".xml")
	if err != nil {

		//fmt.Print(err)
		return false
	}
	err = os.Rename("./storrage/"+GuildID+"/"+IGN+".txt", "./!trash/"+GuildID+"/"+strconv.Itoa(now.UTC().Day())+strconv.Itoa(now.UTC().Hour())+strconv.Itoa(now.UTC().Minute())+strconv.Itoa(now.UTC().Second())+strconv.Itoa(now.UTC().Year())+now.UTC().Weekday().String()+"_old_"+IGN+".txt")
	if err != nil {
		//fmt.Print(err)
		return false
	}

	return true
}

func ReadBlacklist(GuildID string) (blacklisted xml_blacklist_struct, errors bool) {
	filename := "./blacklisted/" + GuildID + Blacklistfilename

	e1 := &xml_blacklist_struct{}
	Check := check_for_file_filename(filename)

	if !Check {
		return *e1, true
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return *e1, true
	}

	err = xml.Unmarshal([]byte(data), &e1)
	if err != nil {
		return *e1, true
	}

	return *e1, false
}
func WriteBlacklist(GuildID string, Blacklisted, Blacklistguilds []string) (errors bool) {
	Check1 := exists_createfolder_for_BL(GuildID)

	if !Check1 {
		//fmt.Println("Err5")
		return false
	}
	filename := "./blacklisted/" + GuildID + Blacklistfilename

	e1 := &xml_blacklist_struct{
		Blacklist:         Blacklisted,
		Blacklistedguilds: Blacklistguilds,
	}

	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		//fmt.Println("Err6")
		return false
	}
	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return false
	}

	return true
}

func Xmlfileconfig_create(GuildID string, g *discordgo.Guild, AdminrolesIDs, AdminIGNs, MemberrolesIDs, FriendIGNs, FriendRoleIDs []string, RegisterchannelID string, Rolecolour, Alliancerolecolour int) (succses bool) {

	filename := filenameConfig(GuildID)

	e1 := &xml_struct_config{
		ServerName:          g.Name,
		ServerID:            g.ID,
		ServerOwnerID:       g.OwnerID,
		AdminRoles:          AdminrolesIDs,
		AdminIGNs:           AdminIGNs,
		GuildMemberRolesIDs: MemberrolesIDs,
		FriendsIngamename:   FriendIGNs,
		FriendRolesIDs:      FriendRoleIDs,
		RegisterchannelID:   RegisterchannelID,
		RoleColour:          Rolecolour,
		AllianceRoleColour:  Alliancerolecolour,
	}

	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return false
	}

	return true
}
func Xmlfileconfig_read(GuildID string) (config xml_struct_config, errors bool) {
	filename := filenameConfig(GuildID)

	e1 := &xml_struct_config{}
	Check := check_for_file_filename(filename)

	if !Check {
		return *e1, true
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return *e1, true
	}

	err = xml.Unmarshal([]byte(data), &e1)
	if err != nil {
		return *e1, true
	}

	return *e1, false
}

/*func main() {

	var b []string
	a, _ := ReadBlacklist("f")

	for _, v := range a.Blacklist {
		if v != "test2" {
			b = append(b, v)
		}
	}

	WriteBlacklist("f", b)

	//fmt.Println(b)
}
*/
