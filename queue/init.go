package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// InitRabbitMQ init the mq connection
func InitRabbitMQ(rabbitURL string, queueName string, enableQuorum bool) (*amqp.Channel, amqp.Queue, error) {
	conn, err := amqp.Dial(rabbitURL + "/")
	if err != nil {
		log.Printf("%s: %s", "Failed to connect to RabbitMQ", err)
		return nil, amqp.Queue{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err)
		return nil, amqp.Queue{}, err
	}

	var queueArgs amqp.Table = nil

	if enableQuorum {
		queueArgs = amqp.Table{
			"x-queue-type": "quorum",
		}
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		queueArgs, // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	return ch, q, nil
}
