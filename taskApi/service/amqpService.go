package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type AmqpService struct {
	channel *amqp.Channel
}

const (
	QueueName = "taskQueue"
)

func NewAmqpService() *AmqpService {
	conn, err := amqp.Dial("amqp://admin:admin217@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return &AmqpService{
		channel: ch,
	}
}

func (as *AmqpService) PublishToQueue(task *model.Task) error {
	message, err := serializeTask(task)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = as.channel.PublishWithContext(ctx, "", QueueName, false, false, amqp.Publishing{
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
