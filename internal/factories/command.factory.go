package factories

import (
	"fmt"
	"github.com/Miuzaki/cukectrl/internal/commands"
	"github.com/Miuzaki/cukectrl/internal/database/repositories"
	"github.com/Miuzaki/cukectrl/pkg/models"
)

type BotInstanceCommand struct {
	br *repositories.BotRepository
}

func NewBotInstanceCommand(br *repositories.BotRepository) *BotInstanceCommand {
	return &BotInstanceCommand{
		br: br,
	}
}

func (f *BotInstanceCommand) Create(event models.Event) (commands.Command, error) {
	switch event.Type {
	case "create":
		return &commands.CreateBot{Event: event, Br: f.br}, nil
	case "delete":
		return &commands.DeleteBot{ID: event.ID.String(), Br: f.br}, nil
	case "start":
		return &commands.StartBot{ID: event.ID.String()}, nil
	case "stop":
		return &commands.StopBot{ID: event.ID.String()}, nil
	default:
		return nil, fmt.Errorf("evento desconhecido: %s", event.Type)
	}
}
