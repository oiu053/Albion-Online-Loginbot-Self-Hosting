package commands

import (
	"github.com/bwmarrin/discordgo"
	operation "github.com/oiu053/Login_Bot_Albion/src/DiscordReaktions"
)

var (
	yearmin     float64 = 1970
	min1        float64 = 1
	monthamount float64 = 12
	maxday      float64 = 31
	Max59       float64 = 59
	Max23       float64 = 23
	Min0        float64 = 0

	Colourmin float64 = 0
	ColourMax float64 = 225
)

func Commandsget() []*discordgo.ApplicationCommand {
	command := []*discordgo.ApplicationCommand{
		{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "timestamp-date",
			Description: "Creates a countdown in (In development)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "yyyy",
					Description: "year",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MinValue:    operation.PointerTo(yearmin),
				},
				{
					Name:        "mm",
					Description: "month",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MinValue:    operation.PointerTo(min1),
					MaxValue:    monthamount,
				},
				{
					Name:        "dd",
					Description: "day",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MinValue:    operation.PointerTo(min1),
					MaxValue:    maxday,
				},
				{
					Name:        "hh",
					Description: "day",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MaxValue:    Max23,
					MinValue:    operation.PointerTo(Min0),
				},
				{
					Name:        "minmin",
					Description: "day",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MaxValue:    Max59,
					MinValue:    operation.PointerTo(Min0),
				},
			},
		},
		{
			Type:        discordgo.ChatApplicationCommand,
			Name:        "countdown",
			Description: "Creates a countdown in ",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "day",
					Description: "days will be added to the momentual timestamp",
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionInteger,
				},
				{
					Name:        "hour",
					Description: "hours will be added to the momentual timestamp",
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MaxValue:    23,
				},
				{
					Name:        "minute",
					Description: "minutes will be added to the momentual timestamp",
					Required:    false,
					Type:        discordgo.ApplicationCommandOptionInteger,
					MaxValue:    59,
				},
			},
		},
		{
			Name:        "help",
			Description: "gives an advise!",
		},
		{
			Name:         "get_false_registrations",
			DMPermission: operation.PointerTo(false),
			Description:  "gets the false registered--> Forceunregisters them!",
		},
		{
			Name:         "alliance-colour-set",
			Description:  "Sets the colour for new created allianceroles! The default is random.",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "r",
					Description: "the Number for the red colour (RGB) => R",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "g",
					Description: "the Number for the red colour (RGB) => G",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "b",
					Description: "the Number for the red colour (RGB) => B",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
			},
		},
		{
			Name:         "alliance-colour-set-hex",
			Description:  "Sets the colour for new created allianceroles! The default is random.",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "hex",
					Description: "the HEX colour",
					Required:    true,
					MinLength:   operation.PointerTo(7),
					MaxLength:   7,
				},
			},
		},
		{
			Name:         "alliance-colour-rand",
			Description:  "Sets the colour for new created allianceroles to random!",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "force-unregister",
			Description:  "force unregisters an IGN from this Server!",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:                     "get-registered-players",
			Description:              "lists everyone who is registered",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionAll)),
		},
		{
			Name:                     "get-unregisteredplayers",
			Description:              "Lists the unregistered players of the Guild by the InGameGuildName",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionAll)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "guildname",
					Description: "Type the Ingame Guildname",
					Required:    true,
				},
			},
		},
		{
			Name:         "guild-colour-set",
			Description:  "Sets the colour for new created Guildroles! The default is random.",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "r",
					Description: "the Number for the red colour (RGB) => R",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "g",
					Description: "the Number for the red colour (RGB) => G",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "b",
					Description: "the Number for the red colour (RGB) => B",
					Required:    true,
					MinValue:    &Colourmin,
					MaxValue:    ColourMax,
				},
			},
		},
		{
			Name:         "guild-colour-set-hex",
			Description:  "Sets the colour for new created Guildroles! The default is random.",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "hex",
					Description: "the HEX colour",
					Required:    true,
					MinLength:   operation.PointerTo(7),
					MaxLength:   7,
				},
			},
		},
		{
			Name:         "guild-colour-rand",
			Description:  "Sets the colour for new created Guildroles to random!",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "clear",
			Description:  "clears the channel but ignores pinned messages!",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "clear-all",
			Description:  "clears the channel!",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "register-channel-set",
			Description:  "sets the registerchannel",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "register-channel-remove",
			Description:  "removes the Registechannel",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:                     "admin_ign-remove",
			Description:              "removes an IGN from the Adminlist",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		}, //askhbakjsbasldjnajsdnkadsa
		{
			Name:                     "admin_ign-add",
			Description:              "adds an IGN to the Adminlist",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		}, //askhbakjsbasldjnajsdnkadsa
		{
			Name:         "check-bl",
			Description:  "checks if a player is blacklisted",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "InGameName",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:                     "adminroles-add",
			Description:              "Adds a role to the Admin roles",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:         "registered_roles-add",
			Description:  "Adds a role to the Member roles",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:                     "friends_roles-add",
			Description:              "Adds a role to the Member roles",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:         "friends_ign-add",
			Description:  "adds an IGN from the Friendslist",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		}, //askhbakjsbasldjnajsdnkadsa
		{
			Name:                     "adminroles-remove",
			Description:              "Removes a role from the Admin roles",
			DMPermission:             operation.PointerTo(false),
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:         "registered_roles-remove",
			Description:  "Removes a role from the Registered roles",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:                     "friends_roles-remove",
			DefaultMemberPermissions: operation.PointerTo(int64(discordgo.PermissionManageServer)),
			Description:              "Removes a role from the Friend roles",
			DMPermission:             operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "role",
					Description: "DiscordRole",
					Type:        discordgo.ApplicationCommandOptionRole,
					Required:    true,
				},
			},
		},
		{
			Name:         "friends_ign-remove",
			Description:  "removes an IGN from the Friendslist",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:         "register",
			Description:  "Registers to this Server!",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:         "unregister",
			Description:  "Unregister from this server",
			DMPermission: operation.PointerTo(false),
		},
		{
			Name:         "blacklist_add",
			Description:  "Blacklists an IGN to this Server!",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:         "blacklist_remove",
			Description:  "Removes an IGN from the Blacklist of this Server!",
			DMPermission: operation.PointerTo(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ign",
					Description: "Ingamename",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
	}

	return command
}

func Testcommands() (testcmd []*discordgo.ApplicationCommand) {

	testcmd = []*discordgo.ApplicationCommand{
		{
			Name:        "guilds",
			Description: "Lists every guild of this bot!",
		},
	}
	return testcmd
}
