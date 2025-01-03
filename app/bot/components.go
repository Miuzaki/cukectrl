package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type ComponentRegistry struct {
	components []*Component
}

var ComponentRegistryGlobal = &ComponentRegistry{}

type Component struct {
	CustomID string
	Handler  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (cr *ComponentRegistry) Register(components ...*Component) {
	cr.components = append(cr.components, components...)
}

func (cr *ComponentRegistry) Init(s *discordgo.Session) {
	log.Info().Msgf("BOT %s COMPONENTS: Registrando componentes...", s.State.User.Username)
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			return
		}

		switch i.Type {
		case discordgo.InteractionMessageComponent:
			{
				customID := i.MessageComponentData().CustomID
				for _, comp := range cr.components {
					if customID == comp.CustomID {
						comp.Handler(s, i)
						return
					}
					log.Warn().Msgf("Erro ao registrar componente '%s'", comp.CustomID)
				}
			}
		case discordgo.InteractionModalSubmit:
			{
				customID := i.ModalSubmitData().CustomID
				for _, comp := range cr.components {
					if customID == comp.CustomID {
						comp.Handler(s, i)
						return
					}
					log.Warn().Msgf("Erro ao registrar componente '%s'", comp.CustomID)
				}
			}
		}
	})
}
