package commands

import (
	"github.com/Miuzaki/cukectrl/app/bot"
	"github.com/Miuzaki/cukectrl/internal/database/repositories"
	"github.com/Miuzaki/cukectrl/internal/instances"
	"github.com/Miuzaki/cukectrl/pkg/models"
	"github.com/Miuzaki/cukectrl/pkg/utils"
)

type CreateBot struct {
	Event models.Event
	Br    *repositories.BotRepository
}

func (c *CreateBot) Execute(manager *instances.InstanceManager) error {
	err := utils.ValidateBotToken(c.Event.Payload)
	if err != nil {
		return err
	}

	err = c.Br.Create(&models.Bot{
		ID:    c.Event.ID,
		Token: c.Event.Payload,
	})

	if err != nil {
		return err
	}

	return manager.AddInstance(c.Event.ID.String(), c.Event.Payload, func(stopChan <-chan struct{}) {
		bot.Init(c.Event.Payload, stopChan)
	})
}
