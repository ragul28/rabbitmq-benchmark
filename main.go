package main

import (
	"flag"
	"fmt"

	"github.com/ragul28/rabbitmq-benchmark/queue"
	"github.com/ragul28/rabbitmq-benchmark/utils"
)

var rabbitURL string
var numWorker int
var role string
var msgSize int

func main() {

	flag.StringVar(&rabbitURL, "url", "amqp://guest:guest@localhost:5672", "Rabbitmq connection string")
	flag.IntVar(&numWorker, "t", 3, "Num of worker threads")
	flag.StringVar(&role, "r", "consumer", "Select consumer or producer")
	flag.IntVar(&msgSize, "s", 10, "producer message size")

	flag.Parse()

	fmt.Println("Runing as", role)

	switch role {
	case "consumer":
		// Start consumer worker threads using goroutine
		for w := 1; w <= numWorker; w++ {
			ch, q := queue.InitRabbitMQ(rabbitURL)
			fmt.Printf("Consumer Worker %d started..\n", w)
			go queue.ConsumerMQ(ch, q)
		}

	case "producer":
		// Start publisher worker threads using goroutine
		for w := 1; w <= numWorker; w++ {
			ch, q := queue.InitRabbitMQ(rabbitURL)
			fmt.Printf("Publisher Worker %d started..\n", w)
			go queue.PublishMQ(ch, q, msgSize)
		}

	default:
		fmt.Printf("Not a valid role, Please choose consumer / publisher")
	}

	utils.CloserHandler()
	// Block main thread to allow running goroutines
	select {}
}
