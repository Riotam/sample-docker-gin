FROM golang:1.17

COPY docker/go/go.mod /go.mod
COPY docker/go/go.sum /go.sum

RUN go mod download

WORKDIR /app

CMD go run cmd/main.go