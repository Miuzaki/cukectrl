package commands

import (
	"github.com/Miuzaki/cukectrl/internal/database/repositories"
	"github.com/Miuzaki/cukectrl/internal/instances"
)

type DeleteBot struct {
	ID string
	Br *repositories.BotRepository
}

func (d *DeleteBot) Execute(manager *instances.InstanceManager) error {

	err := d.Br.Delete(d.ID)
	if err != nil {
		return err
	}

	return manager.DeleteInstance(d.ID)
}
