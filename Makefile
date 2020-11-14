build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor

install:
	go install
	
run:
	go build && ./rabbitmq-benchmark

run_p:
	go build && ./rabbitmq-benchmark -t 1 -r producer -f 1000

run_c:
	go build && ./rabbitmq-benchmark -t 1 -r consumer -debug

start_mq:
	docker-compose up -d

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor
