FROM golang:1.17

COPY docker/go/cmd/app/go.mod /go.mod
COPY docker/go/cmd/app/go.sum /go.sum
WORKDIR /app
RUN ls
RUN go mod download

COPY ../../.. /app

CMD go run main.go