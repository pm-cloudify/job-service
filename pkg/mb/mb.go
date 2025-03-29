package mb

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s # %s\n", msg, err.Error())
	}
}

func InitMessageBroker(addr, user, pass, q_name string) *MessageBroker {
	url := fmt.Sprintf("amqp://%s:%s@%s", user, pass, addr)

	conn, err := amqp.Dial(url)
	failOnError(err, "failed to dial")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		q_name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")

	return &MessageBroker{
		conn, ch, q,
	}
}

func ProduceTextMsg(mb *MessageBroker, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := mb.ch.PublishWithContext(
		ctx,
		"",
		mb.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	failOnError(err, "failed to publish")
	log.Printf("sent msg: %s\n", msg)
}

func AttachConsumer(mb *MessageBroker, handler func(data []byte)) {
	msgs, err := mb.ch.Consume(
		mb.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "consumer register failed")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("reading message!\n")
			handler(d.Body)
		}
	}()

	log.Printf("waiting for message...\n")
	<-forever
}
