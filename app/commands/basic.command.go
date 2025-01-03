package commands

import (
	"github.com/Miuzaki/cukectrl/app/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Basic() *bot.Command {
	return &bot.Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "basic-command",
			Description: "Basic command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "send-message",
					Description: "Send message to chhannel",

					Type:     discordgo.ApplicationCommandOptionChannel,
					Required: true,
				},
			},
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			channelID := options[0].ChannelValue(s)

			_, err := s.ChannelMessageSend(channelID.ID, "oi")

			if err != nil {
				log.Warn().Msgf("Erro ao enviar mensagem no comando 'basic-command' no bot "+s.State.User.Username+", Erro :", err)
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "This is a test embed",
							Description: "This is a test description",
						},
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Yes",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "test",
								},

								discordgo.Button{
									Label:    "I don't know",
									Style:    discordgo.LinkButton,
									Disabled: false,
									URL:      "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🤷",
									},
								},
							},
						},
					},
				},
			})

			if err != nil {
				log.Warn().Msgf("Erro no comando 'basic-command' no bot "+s.State.User.Username+", Erro :", err)
				return
			}

			log.Warn().Msgf("Comando 'basic-command' executado no bot: %s", s.State.User.Username)
		},
	}

}
