FROM golang:latest

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y make

RUN make build

EXPOSE 8080

CMD ["make", "docker-run"]