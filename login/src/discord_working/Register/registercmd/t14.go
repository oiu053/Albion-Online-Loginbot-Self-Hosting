package registercmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	albionAPI "github.com/oiu053/Login_Bot_Albion/src/Albion_Api"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
	xmldb "github.com/oiu053/Login_Bot_Albion/src/XmL_db"
)

var storragefolder string = "storrage/"

func Register(s *discordgo.Session, cmd *discordgo.InteractionCreate) {

	//loading(s, cmd)

	if cmd.Member.User.ID == s.State.User.ID {
		return
	}

	Content1 := cmd.ApplicationCommandData().Options[0].Value

	//fmt.Println(Content1)

	Content := fmt.Sprintf("value: %v", Content1)
	//fmt.Println("t1")
	Content = strings.TrimPrefix(Content, "value: ")
	//Content := strings.TrimPrefix(cmd.Message.Content, cmdreg)

	//fmt.Print(Content)

	Check := xmldb.Check_for_file(cmd.GuildID, cmd.Member.User.ID)

	//fmt.Print(Check)
	if Check {
		//fmt.Println("Error Nr. 1")
		s.ChannelMessageSend(cmd.ChannelID, "You have already registered!")
		return
	}
	Check = xmldb.Check_for_file_txt(cmd.GuildID, Content)
	//fmt.Print(Check)
	if Check {
		s.ChannelMessageSend(cmd.ChannelID, "This acc has already logged in!")
		return
	}

	//check blacklist
	Check = xmldb.Check_for_file_blacklist(cmd.GuildID)
	if Check {
		//fmt.Println("Error Nr. 1")
		a, errors := xmldb.ReadBlacklist(cmd.GuildID)

		if errors {
			s.ChannelMessageSend(cmd.ChannelID, "An Error with the blacklist check had occured. please try again later!")
			return
		}

		Blacklist := a.Blacklist

		for _, Blacklisted_Pl := range Blacklist {
			if strings.ToLower(Blacklisted_Pl) == strings.ToLower(Content) {
				s.ChannelMessageSend(cmd.ChannelID, "You are blacklisted!")
				return
			}
		}
	}

	Guild, Errors := albionAPI.SearchPlayerGuildName(Content)

	if Errors {
		s.ChannelMessageSend(cmd.ChannelID, "Please check the writing of your name!")
		return
	}

	//fmt.Println("t2")

	Alliance, errors2 := albionAPI.SearchPlayerAllianceName(Content)
	if errors2 {
		return
	}

	Check2 := xmldb.Xmlfile_create_auto_cmd(cmd.GuildID, cmd.Member.User.ID, Content, Guild, Alliance, cmd)

	if !Check2 {
		s.ChannelMessageSend(cmd.ChannelID, "Error! Something went wrong")
		return
	}

	config, errors := xmldb.Xmlfileconfig_read(cmd.GuildID)

	if config.RegisterchannelID != "" {
		if config.RegisterchannelID != cmd.ChannelID {
			s.ChannelMessageSend(cmd.ChannelID, "Please use the Registerchannel")
			return
		}
	}

	if errors {
		return
	}

	for _, Role := range config.GuildMemberRolesIDs {
		s.GuildMemberRoleAdd(cmd.GuildID, cmd.Member.User.ID, Role)
	}

	for _, IGN := range config.FriendsIngamename {
		if strings.ToLower(IGN) == strings.ToLower(Content) {
			for _, FrRoles := range config.FriendRolesIDs {
				s.GuildMemberRoleAdd(cmd.GuildID, cmd.Member.User.ID, FrRoles)
			}
		}
	}
	for _, IGN := range config.AdminIGNs {
		if strings.ToLower(IGN) == strings.ToLower(Content) {
			for _, AdRoles := range config.AdminRoles {
				s.GuildMemberRoleAdd(cmd.GuildID, cmd.Member.User.ID, AdRoles)
			}
		}
	}
	//Search Guild
	SearchGuild, err := s.Guild(cmd.GuildID)
	if err != nil {
		//fmt.Println("Err Nr. 3")
		return
	}

	if Guild == "" {
		Guild = "/"
	}

	if Guild == `"GuildName":null` {
		Guild = "/"
	}

	Guildrole, exists := operation.Roleexists(Guild, s, SearchGuild)

	//fmt.Print(Guildrole)
	//fmt.Print(exists)

	if !exists {
		//va := operation.NewGateway(s)

		/*check, errors := xmldb.Xmlfileconfig_read(cmd.GuildID)

		if errors {
			return
		}*/

		var va int64 = 0

		rand.Seed(time.Now().UnixNano())
		RandColourB := rand.Intn(224) + 1
		rand.Seed(time.Now().UnixNano())
		RandColourG := rand.Intn(224) + 1
		rand.Seed(time.Now().UnixNano())
		RandColourR := rand.Intn(224) + 1

		Rechnung1 := RandColourR * 65536
		Rechnung2 := RandColourG * 256
		Rechnung3 := RandColourB

		RandColour := Rechnung1 + Rechnung2 + Rechnung3

		if config.RoleColour != 917039360 {
			RandColour = config.RoleColour
		}

		//fmt.Print(RandColour)

		Guildrole, err = s.GuildRoleCreate(SearchGuild.ID, &discordgo.RoleParams{
			Name:        Guild,
			Color:       operation.PointerTo(RandColour),
			Hoist:       operation.PointerTo(true),
			Permissions: operation.PointerTo(va),
			Mentionable: operation.PointerTo(false),
		})

		if err != nil {
			return
		}

	}

	var Roleordersort []*discordgo.Role

	ROl, err := s.GuildRoles(cmd.GuildID)
	if err != nil {
		return
	}

	for i := range ROl {
		for _, v := range ROl {
			if v.Position == i {
				Roleordersort = append(Roleordersort, v)
			}
		}
	}

	var donotoverwork bool = true
	var NewSort []*discordgo.Role

	for _, v := range Roleordersort {
		if !donotoverwork {
			v.Position = v.Position + 1
			NewSort = append(NewSort, v)
		}
		if donotoverwork {
			var Newroles bool = true
			skipnext := false
			if v.ID == Guildrole.ID {
				skipnext = true
			}
			if !skipnext {
				for _, v2 := range config.GuildMemberRolesIDs {
					if Newroles {
						if v.ID == v2 {
							Newroles = false
							donotoverwork = false
							NewSort = append(NewSort, v)
							newpos := v.Position + 1

							//fmt.Println(v.Position)
							//fmt.Println(newpos)
							Guildrole.Position = newpos
							NewSort = append(NewSort, Guildrole)

						}

					}
				}
				if Newroles {
					NewSort = append(NewSort, v)
				}
			}
		}

	}

	_, err003 := s.GuildRoleReorder(cmd.GuildID, NewSort)

	if err003 != nil {
		s.ChannelMessageSend(cmd.ChannelID, "An error had occured. it wasn't possible to change the position of the Guildrole. Please use /unregister & and reregister again!")
		//fmt.Println(err003)
	}

	s.GuildMemberRoleAdd(cmd.GuildID, cmd.Member.User.ID, Guildrole.ID)

	Nickname := "[" + Guild + "]" + Content

	//alliancerole:
	SearchAlliance, err := s.Guild(cmd.GuildID)

	Alli, errors001 := albionAPI.SearchPlayerAllianceName(Content)

	//fmt.Println(errors001)
	if !errors001 {
		//fmt.Print(Alli)

		if err != nil {
			//fmt.Println("Err Nr. 3")
			return
		}

		alliancerole, exists := operation.Roleexists(Alli, s, SearchAlliance)

		//fmt.Print(Guildrole)
		//fmt.Print(exists)

		if !exists {
			//va := operation.NewGateway(s)

			/*check, errors := xmldb.Xmlfileconfig_read(cmd.GuildID)

			if errors {
				return
			}*/

			var va int64 = 0

			rand.Seed(time.Now().UnixNano())
			RandColourB := rand.Intn(224) + 1
			rand.Seed(time.Now().UnixNano())
			RandColourG := rand.Intn(224) + 1
			rand.Seed(time.Now().UnixNano())
			RandColourR := rand.Intn(224) + 1

			Rechnung1 := RandColourR * 65536
			Rechnung2 := RandColourG * 256
			Rechnung3 := RandColourB

			RandColour := Rechnung1 + Rechnung2 + Rechnung3

			if config.AllianceRoleColour != 917039360 {
				RandColour = config.AllianceRoleColour
			}

			//fmt.Print(RandColour)

			alliancerole, err = s.GuildRoleCreate(SearchAlliance.ID, &discordgo.RoleParams{
				Name:        Alli,
				Color:       operation.PointerTo(RandColour),
				Hoist:       operation.PointerTo(false),
				Permissions: operation.PointerTo(va),
				Mentionable: operation.PointerTo(false),
			})

			if err != nil {
				return
			}

		}

		var Roleordersort []*discordgo.Role

		ROl, err := s.GuildRoles(cmd.GuildID)
		if err != nil {
			return
		}

		for i := range ROl {
			for _, v := range ROl {
				if v.Position == i {
					Roleordersort = append(Roleordersort, v)
				}
			}
		}

		var donotoverwork bool = true
		var NewSort []*discordgo.Role

		for _, v := range Roleordersort {
			if !donotoverwork {
				v.Position = v.Position + 1
				NewSort = append(NewSort, v)
			}
			if donotoverwork {
				var Newroles bool = true
				skipnext := false
				if v.ID == alliancerole.ID {
					skipnext = true
				}
				if !skipnext {
					for _, v2 := range config.GuildMemberRolesIDs {
						if Newroles {
							if v.ID == v2 {
								Newroles = false
								donotoverwork = false
								NewSort = append(NewSort, v)
								newpos := v.Position + 1

								//fmt.Println(v.Position)
								//fmt.Println(newpos)
								alliancerole.Position = newpos
								NewSort = append(NewSort, alliancerole)

							}

						}
					}
					if Newroles {
						NewSort = append(NewSort, v)
					}
				}
			}

		}

		_, err003 := s.GuildRoleReorder(cmd.GuildID, NewSort)

		if err003 != nil {
			s.ChannelMessageSend(cmd.ChannelID, "An error had occured. it wasn't possible to change the position of the Guildrole. Please use /unregister & and reregister again!")
			//fmt.Println(err003)
		}

		s.GuildMemberRoleAdd(cmd.GuildID, cmd.Member.User.ID, alliancerole.ID)
		//fmt.Print(Nickname)
	}

	err = s.GuildMemberNickname(cmd.GuildID, cmd.Member.User.ID, Nickname)

	if err != nil {
		s.ChannelMessageSend(cmd.ChannelID, "An Error had occured. Please check my permissions. You Are already logged In. Please ask an Admin to rename your Nickname.")
		s.ChannelMessageSend(cmd.ChannelID, "You are now logged in!")
		return
	}

	s.ChannelMessageSend(cmd.ChannelID, "You are now logged in!")

	return

}
