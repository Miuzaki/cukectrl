package repositories

import (
	"github.com/Miuzaki/cukectrl/pkg/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type BotRepository struct {
	db *gorm.DB
}

func NewBotRepository(db *gorm.DB) *BotRepository {
	return &BotRepository{db}
}

func (r *BotRepository) Create(bot *models.Bot) error {
	return r.db.Create(bot).Error
}

func (r *BotRepository) Delete(id string) error {
	log.Info().Msgf("BOT REPOSITORY: Deletando bot com id: %s", id)
	return r.db.Where("id = ?", id).Delete(&models.Bot{}).Error
}

func (r *BotRepository) Get(id string) (*models.Bot, error) {
	log.Info().Msgf("BOT REPOSITORY: Procurndo bot com id: %s", id)
	var bot models.Bot
	err := r.db.First(&bot, id).Error
	return &bot, err
}

func (r *BotRepository) GetAll() ([]models.Bot, error) {
	log.Info().Msg("BOT REPOSITORY: Listando bots")
	var bots []models.Bot
	err := r.db.Find(&bots).Error
	return bots, err
}

func (r *BotRepository) Update(bot *models.Bot) error {
	log.Info().Msgf("BOT REPOSITORY: Atualizando bot com id: %s", bot.ID)
	return r.db.Save(bot).Error
}
