package queue

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/abondar24/TaskDistributor/taskApi/model"
	"github.com/abondar24/TaskDistributor/taskData/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"strings"
)

type AmqpProducer struct {
	channel   *amqp.Channel
	queueName *string
}

func NewAmqpProducer(conf *config.Config) *AmqpProducer {

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

	return &AmqpProducer{
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

func (as *AmqpProducer) PublishToQueue(task *model.Task) error {
	message, err := serializeTask(task)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = as.channel.PublishWithContext(ctx, "", *as.queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	})

	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	} else {
		log.Println("Message sent to the RabbitMQ queue")
		return err
	}
}

func serializeTask(task *model.Task) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(task); err != nil {
		log.Println("Error encoding struct:", err)
		return nil, err
	}

	byteSlice := buf.Bytes()
	return byteSlice, nil
}
