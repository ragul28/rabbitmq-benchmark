package main

import (
	"flag"
	"fmt"

	"github.com/ragul28/rabbitmq-benchmark/queue"
	"github.com/ragul28/rabbitmq-benchmark/utils"
)

var rabbitURL string
var numWorker int

func main() {

	flag.StringVar(&rabbitURL, "url", "amqp://guest:guest@localhost:5672", "Rabbitmq connection string")
	flag.IntVar(&numWorker, "t", 3, "Num of worker threads")

	flag.Parse()

	// Start consumer worker threads using goroutine
	for w := 1; w <= numWorker; w++ {
		ch, q := queue.InitRabbitMQ(rabbitURL)
		fmt.Printf("Consumer Worker %d started..\n", w)
		go queue.ConsumerMQ(ch, q)
	}

	utils.CloserHandler()
	// Block main thread to allow running goroutines
	select {}
}
