package commands

import (
	"github.com/Miuzaki/cukectrl/internal/instances"
)

type StopBot struct {
	ID string
}

func (s *StopBot) Execute(manager *instances.InstanceManager) error {
	return manager.StopInstance(s.ID)
}
