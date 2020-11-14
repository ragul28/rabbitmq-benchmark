package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// consumer consume queue messages
func consumer(ch *amqp.Channel, q amqp.Queue, enableQuorum bool) (<-chan amqp.Delivery, error) {

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

// ConsumerMQ worker func
func ConsumerMQ(cfg utils.ConfigStore) {
	ch, q, notify, err := InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
	if err != nil {
		log.Fatal(err)
	}
	msgs, _ := consumer(ch, q, cfg.EnableQuorum)

	// forever := make(chan bool)
	var d amqp.Delivery
	// d := make(chan amqp.Delivery)

	for { //receive loop
		select { //check connection
		case <-notify:
			fmt.Println("Detects connection failuer, retring..")
			ch, q, notify, msgs = failuerRetry(cfg)
		case d = <-msgs:
			if cfg.EnableDebug {
				log.Printf("consumer message: %s\n", d.Body)
			}
		}
	}
			d.Ack(false)
}
	}()

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
