package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

/**
TODO:
1 - consumes reqs from a massage broker
2 - create a job and insert it into database
**/

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s # %s\n", msg, err.Error())
	}
}

func main() {

	conn, err := amqp.Dial("amqp://admin:TestRab1234@localhost:5672")
	failOnError(err, "failed to dial")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")

}
