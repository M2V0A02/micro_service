BINARY_NAME=notification
MAIN_PATH=cmd/service/main.go

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