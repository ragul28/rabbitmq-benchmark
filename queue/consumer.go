package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// ConsumerMQ consume queue messages
func ConsumerMQ(ch *amqp.Channel, q amqp.Queue) {

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// log.Printf("consumer %d message: %s", numWorker, d.Body)
			d.Ack(false)
		}
	}()

	<-forever
}
