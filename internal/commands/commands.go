package commands

import (
	"github.com/Miuzaki/cukectrl/internal/instances"
)

type Command interface {
	Execute(manager *instances.InstanceManager) error
}
