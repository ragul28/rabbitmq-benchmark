package utils

import "flag"

// ConfigStore struct stores the flags
type ConfigStore struct {
	RabbitURL, QueueName, Role          string
	NumWorker, MsgSize, TimeFrequencyMS int
	EnableDebug, EnableQuorum           bool
}

// LoadFlags parse & loads the commandline flags. Returns ConfigStore struct.
func LoadFlags() ConfigStore {

	var cfg ConfigStore

	flag.StringVar(&cfg.RabbitURL, "url", "amqp://guest:guest@localhost:5672", "Rabbitmq connection string")
	flag.StringVar(&cfg.QueueName, "n", "queue", "Select consumer or producer")
	flag.StringVar(&cfg.Role, "r", "consumer", "Select consumer or producer")
	flag.IntVar(&cfg.NumWorker, "t", 3, "Num of worker threads")
	flag.IntVar(&cfg.MsgSize, "s", 10, "producer message size")
	flag.IntVar(&cfg.TimeFrequencyMS, "f", 0, "producer message frequency")
	flag.BoolVar(&cfg.EnableQuorum, "quorum", false, "enable quorum queue type")
	flag.BoolVar(&cfg.EnableDebug, "debug", false, "enable debug logging")

	flag.Parse()

	return cfg
}
