FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd

EXPOSE 8080

RUN apt-get update && apt-get install -y redis-server

CMD sh -c "redis-server --daemonize yes && ./app"
