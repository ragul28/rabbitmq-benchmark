package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// InitRabbitMQ init the mq connection & retunrs channel, queue & error ch
func InitRabbitMQ(rabbitURL string, queueName string, enableQuorum bool) (*amqp.Channel, amqp.Queue, chan *amqp.Error, error) {

	conn, err := amqp.Dial(rabbitURL + "/")
	if err != nil {
		log.Printf("%s: %s", "Failed to connect to RabbitMQ", err)
		return nil, amqp.Queue{}, nil, err
	}

	// connection close notify on error channel
	notify := conn.NotifyClose(make(chan *amqp.Error))

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err)
		return nil, amqp.Queue{}, nil, err
	}

	var queueArgs amqp.Table = nil

	// Enable quorum queue based on quorum flage
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

	return ch, q, notify, nil
}
