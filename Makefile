build:
	@go build -o bin/dockerized-go-web-application

docker-run:
	@./bin/dockerized-go-web-application

test:
	@go test -v ./...