package queue

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ragul28/rabbitmq-benchmark/utils"
	"github.com/streadway/amqp"
)

// publisher Publish messages
func publisher(message string, ch *amqp.Channel, q amqp.Queue, endebug bool) error {

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
		return err
	}

	utils.DebugLogging(fmt.Sprintf("published message: %s\n", message), endebug)

	return nil
}

// PublishMQ worker func
func PublishMQ(cfg utils.ConfigStore) {

	ch, q, _, err := InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}

	for {
		if (ch != nil && q != amqp.Queue{}) {
			rand.Seed(time.Now().UnixNano())
			err = publisher(utils.RandString(cfg.MsgSize), ch, q, cfg.EnableDebug)
			time.Sleep(time.Duration(cfg.TimeFrequencyMS) * time.Millisecond)
		}

		if err != nil {
			// Sleep 15 sec before retry the amqp connection init
			ch, q, _, err = InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
			if err != nil {
				log.Println("Sleep 15 sec before retrying the publish")
				time.Sleep(15 * time.Second)
			}
		}
	}
}
