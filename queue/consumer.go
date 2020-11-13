package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// ConsumerMQ consume queue messages
func ConsumerMQ(ch *amqp.Channel, q amqp.Queue, enableQuorum bool, enableDebug bool) {

	var queueArgs amqp.Table = nil

	if enableQuorum {
		queueArgs = amqp.Table{
			"x-queue-type": "quorum",
		}
	}

	msgs, err := ch.Consume(
		q.Name,    // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		queueArgs, // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			if enableDebug {
				log.Printf("consumer message: %s", d.Body)
			}
			d.Ack(false)
		}
	}()

	<-forever
}
