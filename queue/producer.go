package queue

import (
	"log"
	"math/rand"
	"time"

	"github.com/ragul28/rabbitmq-benchmark/utils"
	"github.com/streadway/amqp"
)

// publisher Publish messages
func publisher(message string, ch *amqp.Channel, q amqp.Queue) {

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}

	// fmt.Println("published message: " + message)
}

// PublishMQ worker func
func PublishMQ(ch *amqp.Channel, q amqp.Queue, msgSize int) {
	for {
		// time.Sleep(10 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		publisher(utils.RandString(msgSize), ch, q)
	}
}
