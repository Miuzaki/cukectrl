package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type CommandRegistry struct {
	commands           []*Command
	registeredCommands []*discordgo.ApplicationCommand
}

var CommandRegistryGlobal = &CommandRegistry{}

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (cr *CommandRegistry) Register(commands ...*Command) {
	cr.commands = append(cr.commands, commands...)
}

func (cr *CommandRegistry) Init(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		for _, cmd := range cr.commands {
			if i.ApplicationCommandData().Name == cmd.Definition.Name {
				cmd.Handler(s, i)
				return
			}
		}
	})
}

func (cr *CommandRegistry) RegisterCommands(s *discordgo.Session) {
	log.Info().Msgf("BOT %s COMMANDS: Registrando comandos...", s.State.User.Username)
	cr.registeredCommands = make([]*discordgo.ApplicationCommand, len(cr.commands))
	for i, v := range cr.commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v.Definition)
		if err != nil {
			log.Error().Err(err).Msgf("BOT %s COMMANDS: Erro ao registrar '%s'", s.State.User.Username, v.Definition.Name)
			continue
		}
		cr.registeredCommands[i] = cmd
		log.Info().Msgf("BOT %s COMMAND: '%s' registrado com sucesso", s.State.User.Username, v.Definition.Name)
	}
}

func (cr *CommandRegistry) DeleteCommands(s *discordgo.Session) {
	log.Printf("BOT %s COMMAND: Deletando comandos...", s.State.User.Username)
	for _, cmd := range cr.registeredCommands {
		if cmd == nil {
			continue
		}
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Error().Err(err).Msgf("BOT %s COMMAND: Erro ao deletar comando '%s'", s.State.User.Username, cmd.Name)

			if err.(*discordgo.RESTError).Response.StatusCode == 401 {
				log.Warn().Msg("BOT COMMAND: Permissão insuficiente para deletar comandos. Verifique se o bot tem permissão de gerenciar comandos.")
				break
			}
			continue
		}
		log.Info().Msgf("BOT %s COMMAND: '%s' deletado com sucesso", s.State.User.Username, cmd.Name)
	}
}
