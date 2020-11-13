package main

import (
	"fmt"

	"github.com/ragul28/rabbitmq-benchmark/queue"
	"github.com/ragul28/rabbitmq-benchmark/utils"
)

func main() {

	cfg := utils.LoadFlags()

	// Start consumer worker threads using goroutine
	for w := 1; w <= cfg.NumWorker; w++ {
		ch, q := queue.InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
		fmt.Printf("Consumer Worker %d started..\n", w)
		go queue.ConsumerMQ(ch, q, cfg.EnableQuorum, cfg.EnableDebug)
	}

	utils.CloserHandler()
	// Block main thread to allow running goroutines
	select {}
}
