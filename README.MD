# Rabbitmq benchmark

Simple rabbitmq benchmark tool which produce high load messages using goroutines as worker threads.  

## Usage 

Benchmark cli needs to be run as consumer & producer in separate shell.

* Running rabbitmq-benchmark as consumer
```sh
./rabbitmq-benchmark -r consumer -wt 4 -url amqp://guest:guest@localhost:5672
```

* Running rabbitmq-benchmark as producer
```sh
./rabbitmq-benchmark -r producer -wt 4 -url amqp://guest:guest@localhost:5672
```

* Basic cli usage
```
Usage of ./rabbitmq-benchmark:
  -r string
        Select consumer or publisher (default "consumer")
  -url string
        Rabbitmq connection string (default "amqp://guest:guest@localhost:5672")
  -wt int
        Num of worker threads (default 3)
```

## Build

Golang build from source.
```sh
git clone github.com/ragul28/rabbitmq-benchmark
cd rabbitmq-benchmark
make build
```

## Rabbitmq Cluster setup

* Run cluster using docker compose
```bash
cd cluster/
docker-compose up -d
```

* Enable classic queue mirroring & auto node sync
```
docker exec -it rabbitmq-01 rabbitmqctl set_policy ha-fed \
    ".*" '{"ha-mode":"all", "federation-upstream-set":"all", "ha-sync-mode":"automatic", }' \
    --priority 1 \
    --apply-to queues
```

### Create Admin user
```bash
# add user
rabbitmqctl add_user admin admin
# set user as admin
rabbitmqctl set_user_tags admin administrator
# set all permission
rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"
```

