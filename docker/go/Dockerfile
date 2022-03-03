FROM golang:1.17

COPY docker/go/app/go.mod /go.mod
COPY docker/go/app/go.sum /go.sum
WORKDIR /app
RUN ls
RUN go mod download

COPY ../.. /app

CMD go run main.go