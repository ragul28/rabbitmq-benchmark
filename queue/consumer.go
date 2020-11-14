package queue

import (
	"fmt"
	"log"
	"time"

	"github.com/ragul28/rabbitmq-benchmark/utils"
	"github.com/streadway/amqp"
)

// consumer consume queue messages
func consumer(ch *amqp.Channel, q amqp.Queue, enableQuorum bool) (<-chan amqp.Delivery, error) {

	var queueArgs amqp.Table = nil

	// Enable quorum queue based on quorum flage
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
		return nil, err
	}

	return msgs, nil
}

// ConsumerMQ worker func
func ConsumerMQ(cfg utils.ConfigStore) {
	ch, q, notify, err := InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}
	msgs, _ := consumer(ch, q, cfg.EnableQuorum)

	var d amqp.Delivery

	for {
		select {
		case <-notify:
			fmt.Println("Detects connection failuer, retring..")
			ch, q, notify, msgs = failuerRetry(cfg)
		case d = <-msgs:
			if cfg.EnableDebug {
				log.Printf("consumer message: %s\n", d.Body)
			}
		}
	}
}

// failuerRetry infinite loop to retry the amqp connection.
func failuerRetry(cfg utils.ConfigStore) (*amqp.Channel, amqp.Queue, chan *amqp.Error, <-chan amqp.Delivery) {
	for {
		ch, q, notify, err := InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
		if err != nil {
			log.Println("Sleep 15 sec before retrying the publish")
			time.Sleep(15 * time.Second)
		} else {
			fmt.Println("Reconnection is successful.")
			msgs, _ := consumer(ch, q, cfg.EnableQuorum)
			return ch, q, notify, msgs
		}
	}
}
