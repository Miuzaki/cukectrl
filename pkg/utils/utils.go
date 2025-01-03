package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Error().Err(err).Msg(msg)
	}
}

func ValidateBotToken(token string) error {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("erro ao criar sessão do Discord: %v", err)
	}
	_, err = session.User("@me")
	if err != nil {
		return fmt.Errorf("token inválido ou erro ao acessar informações do bot: %v", err)
	}

	if session.Identify.Intents == 0 {
		return fmt.Errorf("as intents não estão configuradas. Configure as intents no painel do Discord Developer e no código")
	}
	return nil
}
