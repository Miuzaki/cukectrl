package rabbitmq

import (
	"encoding/json"

	"github.com/Miuzaki/cukectrl/internal/factories"
	"github.com/Miuzaki/cukectrl/internal/instances"
	"github.com/Miuzaki/cukectrl/pkg/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func HandleMessages(msgs <-chan amqp.Delivery, commandFactory *factories.BotInstanceCommand, manager *instances.InstanceManager) {
	for msg := range msgs {
		var event models.Event
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Error().Err(err).Msg("")
			continue
		}

		cmd, err := commandFactory.Create(event)
		if err != nil {
			log.Error().Err(err).Msg("")
			continue
		}

		if err := cmd.Execute(manager); err != nil {
			log.Error().Err(err).Msg("")
		}
	}
}
