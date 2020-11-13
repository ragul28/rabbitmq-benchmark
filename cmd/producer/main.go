package main

import (
	"fmt"

	"github.com/ragul28/rabbitmq-benchmark/queue"
	"github.com/ragul28/rabbitmq-benchmark/utils"
)

func main() {

	cfg := utils.LoadFlags()

	// Start publisher worker threads using goroutine
	for w := 1; w <= cfg.NumWorker; w++ {
		ch, q := queue.InitRabbitMQ(cfg.RabbitURL, cfg.QueueName, cfg.EnableQuorum)
		fmt.Printf("Publisher Worker %d started..\n", w)
		go queue.PublishMQ(ch, q, cfg.MsgSize, cfg.TimeFrequencyMS)
	}

	utils.CloserHandler()
	// Block main thread to allow running goroutines
	select {}
}
