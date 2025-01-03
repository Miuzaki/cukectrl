package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Miuzaki/cukectrl/app/commands"
	"github.com/Miuzaki/cukectrl/app/components"
	"github.com/Miuzaki/cukectrl/internal/database"
	"github.com/Miuzaki/cukectrl/internal/database/repositories"
	"github.com/Miuzaki/cukectrl/internal/factories"
	"github.com/Miuzaki/cukectrl/internal/instances"
	"github.com/Miuzaki/cukectrl/internal/rabbitmq"
	"github.com/Miuzaki/cukectrl/pkg/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	if os.Getenv("DEBUG") == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	db := database.Init()
	commands.Init()
	components.Init()
	botRepository := repositories.NewBotRepository(db)
	manager := instances.NewManager()
	commandFactory := factories.NewBotInstanceCommand(botRepository)

	rabbitMQ := rabbitmq.NewHandler()
	defer rabbitMQ.Close()

	m, err := rabbitMQ.ConsumeMessages()
	utils.FailOnError(err, "Falha ao registrar o consumidor")

	botManager := instances.NewBotManager(manager, botRepository)
	botManager.InitTokensValidator()

	err = botManager.RecoveryAllInstances()
	utils.FailOnError(err, "Falha ao recuperar instâncias")

	go rabbitmq.HandleMessages(m, commandFactory, manager)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Info().Msg("[*] Aguardando mensagens. Para sair, pressione CTRL+C")

	<-ctx.Done()

	log.Info().Msg("[*] Encerrando...")

	if err := manager.StopAllInstances(); err != nil {
		log.Error().Err(err).Msg("Encerrar instâncias")
	} else {
		log.Info().Msg("[*] Todas as instâncias encerradas com sucesso.")
	}
}
