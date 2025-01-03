package components

import (
	"github.com/Miuzaki/cukectrl/app/bot"
)

func Init() {
	bot.ComponentRegistryGlobal.Register(
		Test(),
	)
}
