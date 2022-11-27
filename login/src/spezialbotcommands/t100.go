package spezialbotcommands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func ListGuilds(s *discordgo.Session, cmd *discordgo.InteractionCreate) {
	if cmd.GuildID == "1031132335762059294" {
		if cmd.Member.User.ID == "795936108932366356" {
			for i, v := range s.State.Guilds {
				s.ChannelMessageSendComplex(cmd.ChannelID, &discordgo.MessageSend{
					Content: strconv.Itoa(i+1) + " / " + strconv.Itoa(len(s.State.Guilds)),
					Embed: &discordgo.MessageEmbed{
						Type:        discordgo.EmbedTypeImage,
						Title:       v.Name + ": " + v.ID,
						Description: v.Description,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "Joined at:",
								Value:  v.JoinedAt.String(),
								Inline: false,
							},
							{
								Name:   "VoiceRegion",
								Value:  v.Region,
								Inline: false,
							},
						},
						Color: 255,

						Image: &discordgo.MessageEmbedImage{
							ProxyURL: v.IconURL(),
							URL:      v.IconURL(),
						},

						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL:      v.BannerURL(),
							ProxyURL: v.BannerURL(),
						},
					},
				})
				Inves, err := s.GuildInvites(v.ID)
				if err == nil {
					if len(Inves) != 0 {
						s.ChannelMessageSend(cmd.ChannelID, Inves[0].Guild.Name+": "+Inves[0].Inviter.Username+": https://discord.gg/"+Inves[0].Code)
					}
				}

			}

		}
	}
}
