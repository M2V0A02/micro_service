include makevars.mk

.PHONY: build run test generate compose-up compose-down clean debug

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

test:
	go test ./...

generate:
	go generate ./...

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

clean:
	rm -f $(BINARY_NAME)

debug:
	dlv debug $(MAIN_PATH) --headless --listen=:2435 --api-version=2 --accept-multiclient

wait-db:
	@sleep 5

all: compose-up wait-db run