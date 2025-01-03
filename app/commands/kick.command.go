package commands

import (
	"github.com/Miuzaki/cukectrl/app/bot"
	"github.com/Miuzaki/cukectrl/app/bot/builders"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Kick() *bot.Command {
	return &bot.Command{
		Definition: &discordgo.ApplicationCommand{
			Name:                     "kick",
			Description:              "Kick a member of server",
			DefaultMemberPermissions: &bot.KickPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "target",
					Description: "The member you want to kick",

					Type:     discordgo.ApplicationCommandOptionMentionable,
					Required: true,
				},
			},
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			user := options[0].UserValue(s)

			guild, err := s.Guild(i.GuildID)

			if err != nil {
				log.Warn().Msgf("Erro ao pegar guild no comando 'kick' no bot "+s.State.User.Username+", Erro :", err)
				return
			}

			err = s.GuildMemberDeleteWithReason(guild.ID, user.ID, "Kicked by "+i.Member.User.Username)

			if err != nil {
				log.Warn().Msgf("Erro ao kickar membro no comando 'kick' no bot "+s.State.User.Username+", Erro :", err)
				return
			}
			embed := builders.NewEmbed().SetColor(0xff0000).SetTitle("Membro Kickado").SetDescription("O membro " + user.Mention() + " foi kickado por " + i.Member.Mention()).Build()
			response := builders.CreateEmbedResponse(embed)
			err = s.InteractionRespond(i.Interaction, response)

			if err != nil {
				log.Warn().Msgf("Erro ao responder interação no comando 'kick' no bot "+s.State.User.Username+", Erro :", err)
				return
			}

		},
	}

}
