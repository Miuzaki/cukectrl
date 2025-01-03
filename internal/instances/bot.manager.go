package instances

import (
	"github.com/Miuzaki/cukectrl/app/bot"
	"github.com/Miuzaki/cukectrl/internal/database/repositories"
	"github.com/Miuzaki/cukectrl/pkg/utils"
	"github.com/rs/zerolog/log"

	"time"
)

type BotManager struct {
	m  *InstanceManager
	br *repositories.BotRepository
}

func NewBotManager(manager *InstanceManager, repository *repositories.BotRepository) *BotManager {
	return &BotManager{m: manager, br: repository}
}

func (bm *BotManager) RecoveryAllInstances() error {
	log.Info().Msg("RECOVERY: Iniciando recuperação de instâncias")
	allBots, err := bm.br.GetAll()

	if err != nil {
		return err
	}

	for _, b := range allBots {
		err := bm.m.AddInstance(b.ID.String(), b.Token, func(stopChan <-chan struct{}) {
			err := utils.ValidateBotToken(b.Token)
			if err != nil {
				log.Error().Err(err).Msgf("Token inválido")

				err := bm.br.Delete(b.ID.String())
				if err != nil {
					log.Error().Err(err).Msgf("Erro ao deletar bot")
				}

				log.Warn().Msgf("RECOVERY: Bot deletado por token invalido: %s", b.ID.String())
				return
			}
			bot.Init(b.Token, stopChan)
		})

		if err != nil {
			log.Error().Err(err).Msgf("Erro ao criar bot")
			continue
		}

		if err := bm.m.StartInstance(b.ID.String()); err != nil {
			log.Error().Err(err).Msgf("Erro ao iniciar instância recuperada")
		}

		log.Info().Msgf("Instância recuperada: %s", b.ID.String())
	}
	return nil
}

func (bm *BotManager) InitTokensValidator() {
	err := bm.m.AddInstance("discord-instances-manager", "", func(stopChan <-chan struct{}) {
		bm.validatorTicker()
	})
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao adicionar instância")
	}

	if err = bm.m.StartInstance("discord-instances-manager"); err != nil {
		log.Error().Err(err).Msgf("Erro ao iniciar a instância")
	}
}

func (bm *BotManager) validatorTicker() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		bm.processValidator()
	}
}

func (bm *BotManager) processValidator() {
	for _, instance := range bm.m.GetAllInstances() {
		if instance.GetID() == "discord-getAllInstances-manager" || instance.Reference == "" {
			continue
		}

		if err := bm.validateAndManage(instance); err != nil {
			log.Error().Err(err).Msgf("Erro ao gerenciar instância")
		}
	}
}

func (bm *BotManager) validateAndManage(instance *Instance) error {
	if err := utils.ValidateBotToken(instance.Reference); err != nil {
		if instance.IsRunning {
			if stopErr := instance.Stop(); stopErr != nil {
				log.Error().Err(err).Msgf("Erro ao parar a instância %s: %s", instance.GetID(), stopErr)
			}
		}
		return bm.m.DeleteInstance(instance.GetID())
	}
	return nil
}
