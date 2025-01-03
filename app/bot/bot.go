package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Init(token string, stopChan <-chan struct{}) {
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Error().Err(err).Msgf("BOT %s: Erro ao iniciar bot do Discord", token)
		return
	}

	// defer func() {
	// 	err := sess.Close()
	// 	if err != nil {
	// 		log.Info().Msg("Erro ao encerrar sessão do bot: %s", err)
	// 	}
	// }()

	err = sess.Open()
	if err != nil {
		log.Error().Err(err).Msgf("BOT %s: Erro ao abrir conexão com o Discord", sess.State.User.Username)
		return
	}

	CommandRegistryGlobal.RegisterCommands(sess)
	CommandRegistryGlobal.Init(sess)
	ComponentRegistryGlobal.Init(sess)
	log.Info().Msgf("BOT %s: Iniciado com sucesso", sess.State.User.Username)

	<-stopChan
	log.Info().Msgf("BOT %s: Recebido sinal para encerrar", sess.State.User.Username)
	CommandRegistryGlobal.DeleteCommands(sess)
}
