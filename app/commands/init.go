package commands

import (
	"github.com/Miuzaki/cukectrl/app/bot"
)

func Init() {
	bot.CommandRegistryGlobal.Register(
		Basic(),
		Kick(),
	)
}
