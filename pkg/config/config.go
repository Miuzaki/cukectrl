package config

import (
	"errors"
	"os"
)

// Config estrutura para armazenar configurações do aplicativo
type Config struct {
	RabbitMQURL   string
	RabbitMQQueue string
}

// LoadConfig carrega e valida as variáveis de ambiente necessárias
func LoadConfig() (*Config, error) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("RABBITMQ_QUEUE")

	if rabbitURL == "" {
		return nil, errors.New("RABBITMQ_URL não está definido")
	}
	if queueName == "" {
		return nil, errors.New("RABBITMQ_QUEUE não está definido")
	}

	return &Config{
		RabbitMQURL:   rabbitURL,
		RabbitMQQueue: queueName,
	}, nil
}
