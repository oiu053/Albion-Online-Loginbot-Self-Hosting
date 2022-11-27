package main

//Rolecolour Implementieren!!!! (var cammands & register & setclour[R: int 1-225],[G: int 1-225],[B: int 1-225], rand_colour/default int = 0)

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	operations "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
	commands2 "github.com/oiu053/Login_Bot_Albion/src/commands"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Admin/IGN/adminignadd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Admin/IGN/adminignremove"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Admin/Role/adminroleadd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Admin/Role/adminrolerem"
	alliancecolourrand "github.com/oiu053/Login_Bot_Albion/src/discord_working/Alliance-colour/rand"
	alicolourset "github.com/oiu053/Login_Bot_Albion/src/discord_working/Alliance-colour/set"
	alicolourhex "github.com/oiu053/Login_Bot_Albion/src/discord_working/Alliance-colour/setHex"
	friendignadd "github.com/oiu053/Login_Bot_Albion/src/discord_working/Friends/FriendsIGN/add"
	friendignremove "github.com/oiu053/Login_Bot_Albion/src/discord_working/Friends/FriendsIGN/remove"
	friroleadd "github.com/oiu053/Login_Bot_Albion/src/discord_working/Friends/Roles/friroladd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Friends/Roles/frirolrem"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/Role/registerroleadd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/Role/registerrolerem"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/forceunreg"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/registercmd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/registerechannel/registerchanneladd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/registerechannel/registerchannelrem"
	getregpers "github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/registration/getregisteredpersons"
	getunregpers "github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/registration/grtunregpers"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/Register/unregistercmd"
	autocountdown "github.com/oiu053/Login_Bot_Albion/src/discord_working/auto-countdown"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/blacklist/blacklistadd"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/blacklist/blacklistrem"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/blacklist/checkbl"
	check1 "github.com/oiu053/Login_Bot_Albion/src/discord_working/check-f-wrong-registrations"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/clear"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/getfalseregistered"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/getfalseregistered/funcforceunregisterbuttonanswer"
	gcolourrand "github.com/oiu053/Login_Bot_Albion/src/discord_working/guildcolour/rand"
	gcolourset "github.com/oiu053/Login_Bot_Albion/src/discord_working/guildcolour/set"
	gcolourhex "github.com/oiu053/Login_Bot_Albion/src/discord_working/guildcolour/sethash"
	"github.com/oiu053/Login_Bot_Albion/src/discord_working/timestampcreator"
	env "github.com/oiu053/Login_Bot_Albion/src/env"
	"github.com/oiu053/Login_Bot_Albion/src/spezialbotcommands"
)

// var storragefolder string = "storrage/"
var Colourmin float64 = 0
var ColourMax float64 = 225

var commands = commands2.Commandsget()
var commandstest = commands2.Testcommands()

func NewJoin(_ *discordgo.Session, Guild *discordgo.GuildCreate) {

	filename := "./storrage/" + Guild.ID + ".xml"

	exists := check_for_file_filename(filename)

	Blacklisted := []string{}

	xmldb.WriteBlacklist(Guild.ID, Blacklisted, Blacklisted)

	//fmt.Print(exists)

	if exists {
		//fmt.Print(exists)
	} else {
		if !exists {
			xmldb.Xmlfileconfig_create(Guild.ID, Guild.Guild, []string{}, []string{}, []string{}, []string{}, []string{}, "", 917039360, 0)
			//fmt.Print("Fatalerror")
		}
	}
}
func NewLeaf(_ *discordgo.Session, Guild *discordgo.GuildDelete) {
	filename := "./storrage/" + Guild.ID + ".xml"

	exists := check_for_file_filename(filename)

	//fmt.Print(exists)

	if exists {
		xmldb.DeleatefilesoftheGuild(Guild.ID)
	}
}
func main() {

	os.Environ()
	Token := env.Token()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		//fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	//dg.AddHandler(messageCreate)
	dg.AddHandlerOnce(ready)
	dg.AddHandler(reconnect)
	dg.AddHandler(commandHandler)
	dg.AddHandler(NewJoin)
	dg.AddHandler(NewLeaf)
	dg.AddHandler(delMessages)
	dg.AddHandler(MessageUpdate)
	dg.AddHandler(MessageDel)
	//dg.AddHandler(delcommands)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	commandsNow, err := dg.ApplicationCommandBulkOverwrite(env.ApplicationID(), "", commands)
	if err != nil {
		fmt.Println(err)
	}

	commandstest, err := dg.ApplicationCommandBulkOverwrite(env.ApplicationID(), "1031132335762059294", commandstest)
	if err != nil {
		fmt.Println(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Print("This bot had been created by the AlbionOnline Player Oui053. Had this bot been edited by someone????? \n I Hope this bot can help you. Ps the Orriginal replit is https://github.com/oiu053/Albion-Online-Loginbot-Self-Hosting")
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	delcommands(dg, commandsNow)
	// Cleanly close down the Discord session.

	fmt.Print(" \n Closed!")

	delcommandstest(dg, commandstest)
	dg.Close()
}

func reconnect(s *discordgo.Session, _ *discordgo.Resumed) {
	//fmt.Print("reconnected/ reload files")
	serverJoin(s)
}
func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Ready ")
	fmt.Println(" Logged in as " + r.User.Username + "#" + r.User.Discriminator)

	/*_, err := s.ApplicationCommandBulkOverwrite(env.ApplicationID(), "", commands)

	if err != nil {
		//fmt.Print(err)
	}*/
	serverJoin(s)

}

func commandHandler(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//readchatlog.Slashcommandusuage(s, cmd)

	for _, IDcheck := range env.BlacklistedGuildIDs() {
		if IDcheck == cmd.GuildID {
			operations.Commandanswer(s, cmd, "Your Discord Server/Guild has been disabled after a Request!")
			return
		}

	}

	if cmd.Interaction.Type == discordgo.InteractionMessageComponent {
		if cmd.MessageComponentData().CustomID == "forceunregister" {
			funcforceunregisterbuttonanswer.Forceunregister(s, cmd)
		}
	}
	//serverJoin(s)
	//testcmd

	if cmd.Interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	operations.Commandanswer(s, cmd, "_ _")

	if cmd.ApplicationCommandData().Name == "guilds" {
		spezialbotcommands.ListGuilds(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "countdown" {
		autocountdown.Autocontent(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "timestamp-date" {
		timestampcreator.Timestamp(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "get_false_registrations" {
		getfalseregistered.Get_false_registered(s, cmd)
	}
	//normalcmd

	if cmd.ApplicationCommandData().Name == "check-f-wrong-registrations" {

		check1.Check_f_wrong_registrations(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "alliance-colour-set-hex" {
		alicolourhex.Alliance_colour_set_hex(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "alliance-colour-set" {
		alicolourset.Alliance_colour_set(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "alliance-colour-rand" {
		alliancecolourrand.Alliancecolourrand(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "get-registered-players" {
		getregpers.Get_registered_players(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "get-unregisteredplayers" {
		getunregpers.Get_unregisteredplayers(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "guild-colour-set-hex" {
		gcolourhex.Guildcoloursethex(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "guild-colour-set" {
		gcolourset.Guildcolourset(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "guild-colour-rand" {
		gcolourrand.Guildcolourrand(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "clear" {
		clear.Clear1(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "clear-all" {
		clear.Clear_all2(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "check-bl" {
		checkbl.Checkbl(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "friends_ign-remove" {
		friendignremove.Friends_ign_remove(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "friends_ign-add" {
		friendignadd.Friends_ign_add(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "admin_ign-add" {
		adminignadd.Admin_ign_add(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "admin_ign-remove" {
		adminignremove.Admin_ign_remove(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "friends_roles-add" {
		friroleadd.Friends_roles_add(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "registered_roles-add" {
		registerroleadd.Registered_roles_add(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "adminroles-add" {
		adminroleadd.Adminroles_add(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "friends_roles-remove" {
		frirolrem.Friends_roles_remove(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "registered_roles-remove" {
		registerrolerem.Registered_roles_remove(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "adminroles-remove" {
		adminrolerem.Adminroles_remove(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "register" {
		registercmd.Register(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "unregister" {
		unregistercmd.Unregister(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "blacklist_add" {
		blacklistadd.Blacklist(s, cmd)
	}
	if cmd.ApplicationCommandData().Name == "force-unregister" {
		forceunreg.Force_unregister(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "blacklist_remove" {
		blacklistrem.Blacklist_remove(s, cmd)
	}

	//jhdvakhfdkahgdfkaskd
	if cmd.ApplicationCommandData().Name == "register-channel-set" {
		registerchanneladd.Registerchannelset(s, cmd)
	}

	if cmd.ApplicationCommandData().Name == "register-channel-remove" {
		registerchannelrem.Registerchannelremove(s, cmd)
	}
}

func MessageUpdate(s *discordgo.Session, msg *discordgo.MessageUpdate) {
	readchatlog.Chatmessagedetected_edite(s, msg)
}
func MessageDel(s *discordgo.Session, msg *discordgo.MessageDelete) {
	readchatlog.Chatmessagedetected_del(s, msg)
}

func serverJoin(s *discordgo.Session) {
	Guilds := s.State.Guilds

	for _, Guild := range Guilds {
		filename := "./storrage/" + Guild.ID + ".xml"
		exists := check_for_file_filename(filename)

		//fmt.Print(exists)

		if exists {
			//fmt.Print(exists)
		} else {
			if !exists {
				Guildstatement, err := s.Guild(Guild.ID)
				if err != nil {
					return
				}
				xmldb.Xmlfileconfig_create(Guild.ID, Guildstatement, []string{}, []string{}, []string{}, []string{}, []string{}, "", 917039360, 0)
				//fmt.Print("Fatalerror")
			}
		}

		exists = xmldb.Check_for_file_blacklist(Guild.ID)

		if exists {
			//fmt.Print(exists)
		} else {
			if !exists {
				xmldb.WriteBlacklist(Guild.ID, []string{}, []string{})
				//fmt.Print("Fatalerror")
			}
		}
	}
}

func check_for_file_filename(filename string) (exists bool) {

	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}

func delMessages(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == s.State.User.ID {

		if msg.Type == discordgo.MessageTypeChannelPinnedMessage {
			s.ChannelMessageDelete(msg.ChannelID, msg.ID)
			return
		}

		timer1 := time.NewTimer(60 * time.Second)
		<-timer1.C
		getmessage, err := s.ChannelMessage(msg.ChannelID, msg.Message.ID)
		if err != nil {
			return
		}
		if getmessage.Pinned {
			return
		}
		s.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}
	//fmt.Print("msg")

}

func delcommands(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) {

	if env.Delcommands() {
		//fmt.Print("Removing commands...")
		for _, cmd := range cmds {
			s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		}
	}

}
func delcommandstest(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) {

	//fmt.Print("Removing commands...")
	for _, cmd := range cmds {
		s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
	}

}
