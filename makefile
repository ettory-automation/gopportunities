.PHONY: default run build test docs clean
APP_NAME=gopportunities

default: run

docs:
	@swag init --parseDependency --parseInternal

run: docs
	@go run main.go

build: docs
	@go build -o $(APP_NAME) main.go

test:
	@go test ./...

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
