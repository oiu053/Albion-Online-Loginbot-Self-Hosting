package readchatlog

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	operations "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

func filenameLog(GuildID, ChannelID string) (filename string) {

	a := operations.Folder_exists(Logfolder + GuildID)
	if !a {
		operations.Create_Folder(Logfolder + GuildID)
	}
	return Logfolder + GuildID + "/" + ChannelID + ".txt"
}

var Logfolder string = "log/"

func AppendFile(_ *discordgo.Session, msg *discordgo.MessageCreate) {

	filepath := filenameLog(msg.GuildID, msg.ChannelID)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//time.Now()
	var timestamp string = time.Now().UTC().String()

	kind_of_message := "Written-Message!"

	var basemessage string = "\n----------------------------------------\n\n" + kind_of_message + "\nTimestamp:\n	" + timestamp + "\nAuther: \n	" + msg.Author.Username + ":\nAuthorID: \n	" + msg.Author.ID + "\nMessageID:\n	" + msg.Message.ID + "\nChannelID:\n	" + msg.ChannelID + "\nContentent:"

	addmessage := basemessage
	for _, v := range strings.Split(msg.Content, "\n") {
		addmessage = addmessage + "\n	" + v
	}
	_, err = file.WriteString(addmessage)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("\nLength: %d bytes", len)
	//fmt.Printf("\nFile Name: %s", file.Name())
}

func Createfile(s *discordgo.Session, cmd *discordgo.MessageCreate) {
	filepath := filenameLog(cmd.GuildID, cmd.ChannelID)

	guild, err := s.Guild(cmd.GuildID)

	if err != nil {
		return
	}

	channel, err := s.Channel(cmd.ChannelID)
	if err != nil {
		return
	}

	data := []byte(cmd.GuildID + "\n" + guild.Name + "\n" + guild.Description + "\n\n" + guild.OwnerID + "\n" + guild.IconURL() + "\n\n" + channel.Name + "\n" + channel.Topic)

	os.WriteFile(filepath, data, os.ModePerm)

	exists := xmldb.Check_for_file_Guildlog_txt(cmd.GuildID, cmd.ChannelID)

	if exists {
		AppendFile(s, cmd)
	} else {
		return
	}
}

func Chatmessagedetected(s *discordgo.Session, msg *discordgo.MessageCreate) {

	filenameLog(msg.GuildID, msg.ChannelID)
	exists := xmldb.Check_for_file_Guildlog_txt(msg.GuildID, msg.ChannelID)

	if exists {
		AppendFile(s, msg)
	} else {
		Createfile(s, msg)
	}
}

func Chatmessagedetected_edite(s *discordgo.Session, msg *discordgo.MessageUpdate) {
	filenameLog(msg.GuildID, msg.ChannelID)
	exists := xmldb.Check_for_file_Guildlog_txt(msg.GuildID, msg.ChannelID)

	if exists {
		AppendFile_EditeMessage(s, msg)
	} else {
		Createfile_editemessage(s, msg)
	}
}
func Chatmessagedetected_del(s *discordgo.Session, msg *discordgo.MessageDelete) {
	filenameLog(msg.GuildID, msg.ChannelID)
	exists := xmldb.Check_for_file_Guildlog_txt(msg.GuildID, msg.ChannelID)

	if exists {
		AppendFile_DelMessage(s, msg)
	} else {
		Createfile_Delmessage(s, msg)
	}
}

// func EditeMessage()
func AppendFile_EditeMessage(_ *discordgo.Session, msg *discordgo.MessageUpdate) {

	filepath := filenameLog(msg.GuildID, msg.ChannelID)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//time.Now()
	var timestamp string = time.Now().UTC().String()

	kind_of_message := "Edit-Message!"

	var basemessage string = "\n----------------------------------------\n\n" + kind_of_message + "\nTimestamp:\n	" + timestamp + "\nAuther: \n	" + msg.Author.Username + ":\nAuthorID: \n	" + msg.Author.ID + "\nMessageID:\n	" + msg.Message.ID + "\nChannelID:\n	" + msg.ChannelID + "\nContentent:"

	addmessage := basemessage
	for _, v := range strings.Split(msg.Content, "\n") {
		addmessage = addmessage + "\n	" + v
	}
	_, err = file.WriteString(addmessage)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("\nLength: %d bytes", len)
	//fmt.Printf("\nFile Name: %s", file.Name())
}

func Createfile_editemessage(s *discordgo.Session, cmd *discordgo.MessageUpdate) {
	filepath := filenameLog(cmd.GuildID, cmd.Message.ID)

	guild, err := s.Guild(cmd.GuildID)

	if err != nil {
		return
	}

	channel, err := s.Channel(cmd.ChannelID)
	if err != nil {
		return
	}

	data := []byte(cmd.GuildID + "\n" + guild.Name + "\n" + guild.Description + "\n\n" + guild.OwnerID + "\n" + guild.IconURL() + "\n\n" + channel.Name + "\n" + channel.Topic)

	os.WriteFile(filepath, data, os.ModePerm)

	exists := xmldb.Check_for_file_Guildlog_txt(cmd.GuildID, cmd.ChannelID)

	if exists {
		AppendFile_EditeMessage(s, cmd)
	} else {
		return
	}
}

func AppendFile_DelMessage(_ *discordgo.Session, msg *discordgo.MessageDelete) {

	filepath := filenameLog(msg.GuildID, msg.ChannelID)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//time.Now()
	var timestamp string = time.Now().UTC().String()

	kind_of_message := "Delete-Message!"

	var basemessage string = "\n----------------------------------------\n\n" + kind_of_message + "\nTimestamp:\n	" + timestamp + "\nMessageID:\n	" + msg.Message.ID + "\nChannelID:\n	" + msg.ChannelID

	addmessage := basemessage

	_, err = file.WriteString(addmessage)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("\nLength: %d bytes", len)
	//fmt.Printf("\nFile Name: %s", file.Name())
}

func Createfile_Delmessage(s *discordgo.Session, cmd *discordgo.MessageDelete) {
	filepath := filenameLog(cmd.GuildID, cmd.ChannelID)

	guild, err := s.Guild(cmd.GuildID)

	if err != nil {
		return
	}

	channel, err := s.Channel(cmd.ChannelID)
	if err != nil {
		return
	}

	data := []byte(cmd.GuildID + "\n" + guild.Name + "\n" + guild.Description + "\n\n" + guild.OwnerID + "\n" + guild.IconURL() + "\n\n" + channel.Name + "\n" + channel.Topic)

	os.WriteFile(filepath, data, os.ModePerm)

	exists := xmldb.Check_for_file_Guildlog_txt(cmd.GuildID, cmd.ChannelID)

	if exists {
		AppendFile_DelMessage(s, cmd)
	} else {
		return
	}
}

func Slashcommandusuage(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	filenameLog(cmd.GuildID, cmd.ChannelID)
	exists := xmldb.Check_for_file_Guildlog_txt(cmd.GuildID, cmd.ChannelID)

	if exists {
		AppendFilecmd(s, cmd)
	} else {
		Createfilecmd(s, cmd)
	}
}

func AppendFilecmd(_ *discordgo.Session, msg *discordgo.InteractionCreate) {

	filepath := filenameLog(msg.GuildID, msg.ChannelID)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//time.Now()
	var timestamp string = time.Now().UTC().String()

	kind_of_message := "Written-Message!"

	var addmessage string = "\n----------------------------------------\n\n" + kind_of_message + "\nTimestamp:\n	" + timestamp + "\nAuther: \n	" + msg.Member.Nick + ":\nAuthorID: \n	" + msg.Member.User.ID + "\nMessageID:\n	" + msg.Message.ID + "\nChannelID:\n	" + msg.ChannelID + "\nContentent:"

	//for _, v := range strings.Split(msg.Content, "\n") {
	addmessage = addmessage + "\n	" + msg.Interaction.Data.Type().String() + "\n" + msg.ApplicationCommandData().Name + "\n"
	//}

	/*for _, v := range msg.ApplicationCommandData().Options {
		addmessage = addmessage + "\n" + v.Name + ": " + fmt.Sprintf("%v", v.Value)
	}

	addmessage = addmessage + "\n"

	for _, v := range msg.MessageComponentData().Values {
		addmessage = addmessage + "\n" + v
	}*/
	_, err = file.WriteString(addmessage)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("\nLength: %d bytes", len)
	//fmt.Printf("\nFile Name: %s", file.Name())
}
func Createfilecmd(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	filepath := filenameLog(cmd.GuildID, cmd.ChannelID)

	guild, err := s.Guild(cmd.GuildID)

	if err != nil {
		return
	}

	channel, err := s.Channel(cmd.ChannelID)
	if err != nil {
		return
	}

	data := []byte(cmd.GuildID + "\n" + guild.Name + "\n" + guild.Description + "\n\n" + guild.OwnerID + "\n" + guild.IconURL() + "\n\n" + channel.Name + "\n" + channel.Topic)

	os.WriteFile(filepath, data, os.ModePerm)

	exists := xmldb.Check_for_file_Guildlog_txt(cmd.GuildID, cmd.ChannelID)

	if exists {
		AppendFilecmd(s, cmd)
	} else {
		return
	}
}
