build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w"

build_mac:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -ldflags="-s -w"

install:
	go install
	
run:
	go build && ./rabbitmq-benchmark

run_p:
	go build && ./rabbitmq-benchmark -t 2 -r producer -debug -f 1000

run_c:
	go build && ./rabbitmq-benchmark -t 5 -r consumer -debug

start_mq:
	docker-compose up -d

mod:
	go mod tidy
	go mod verify
	go mod vendor
