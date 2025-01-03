package commands

import (
	"github.com/Miuzaki/cukectrl/internal/instances"
	"github.com/Miuzaki/cukectrl/pkg/utils"
)

type StartBot struct {
	ID string
}

func (s *StartBot) Execute(manager *instances.InstanceManager) error {
	i, err := manager.GetInstance(s.ID)

	if err != nil {
		return err
	}

	err = utils.ValidateBotToken(i.Reference)
	if err != nil {
		if !i.IsRunning {
			_ = manager.StopInstance(s.ID)
		}
		_ = manager.DeleteInstance(s.ID)
		return err
	}
	return manager.StartInstance(s.ID)
}
