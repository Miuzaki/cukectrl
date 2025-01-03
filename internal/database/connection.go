package database

import (
	"fmt"
	"os"

	"github.com/Miuzaki/cukectrl/pkg/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	log.Info().Msg("DATABASE: Iniciando conexão com o banco de dados")

	localhost := os.Getenv("POSTGRES_HOST")

	user := os.Getenv("POSTGRES_USER")

	database := os.Getenv("POSTGRES_DB")

	password := os.Getenv("POSTGRES_PASSWORD")

	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", localhost, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Bot{})
	log.Info().Msg("DATABASE: Migrando tabelas")
	if err != nil {
		panic(err)
	}

	log.Info().Msg("Conexão com o banco de dados estabelecida")
	return db
}
