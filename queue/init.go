package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// InitRabbitMQ init the mq connection
func InitRabbitMQ(rabbitURL string) (*amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial(rabbitURL + "/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"publisher", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	return ch, q
}
