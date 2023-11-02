package queue

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"strings"
)

type AmqpConsumer struct {
	channel   *amqp.Channel
	queueName *string
}

func NewAmqpConsumer(conf *config.Config) *AmqpConsumer {

	conn, err := amqp.Dial(buildConnectionUri(conf))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		conf.Broker.QueueName, // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	return &AmqpConsumer{
		channel:   ch,
		queueName: &conf.Broker.QueueName,
	}
}

func buildConnectionUri(conf *config.Config) string {
	var uri strings.Builder

	uri.WriteString("amqp://")
	uri.WriteString(conf.Broker.Username)
	uri.WriteString(":")
	uri.WriteString(conf.Broker.Password)
	uri.WriteString("@")
	uri.WriteString(conf.Broker.Host)
	uri.WriteString(":")
	uri.WriteString(strconv.Itoa(conf.Broker.Port))
	uri.WriteString("/")

	return uri.String()
}

func (al *AmqpConsumer) ReadFromQueue() {
	msgs, err := al.channel.Consume(
		*al.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	} else {
		log.Println("Consumer registered successfully")
	}

	for d := range msgs {
		log.Printf("Received a message: %s\n", d.Body)
	}
}
