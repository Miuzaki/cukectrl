package rabbitmq

import (
	"github.com/Miuzaki/cukectrl/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"

	"os"
)

type Handler struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName string
}

func NewHandler() *Handler {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	utils.FailOnError(err, "RABBITMQ HANDLER: Falha ao conectar ao RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "RABBITMQ HANDLER: Falha ao abrir um canal")

	queueName := os.Getenv("RABBITMQ_QUEUE")
	_, err = ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	utils.FailOnError(err, "RABBITMQ HANDLER: Falha ao declarar uma fila")

	return &Handler{
		conn:      conn,
		ch:        ch,
		queueName: queueName,
	}
}

func (r *Handler) Close() {
	if err := r.ch.Close(); err != nil {
		log.Error().Err(err).Msg("RABBITMQ HANDLER: Erro ao fechar o canal")
	}
	if err := r.conn.Close(); err != nil {
		log.Error().Err(err).Msg("RABBITMQ HANDLER: Erro ao fechar a conexão")
	}
}

func (r *Handler) ConsumeMessages() (<-chan amqp.Delivery, error) {
	return r.ch.Consume(
		r.queueName, "", true, false, false, false, nil,
	)
}
