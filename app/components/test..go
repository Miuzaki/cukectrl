package components

import (
	"github.com/Miuzaki/cukectrl/app/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Test() *bot.Component {
	return &bot.Component{
		CustomID: "test",
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Você clicou no botão de exemplo!",
				},
			})
			if err != nil {
				log.Warn().Msgf("Erro ao responder interação: %v", err)
			}
		},
	}
}
