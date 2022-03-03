FROM golang:1.17

ENV GO111MODULE="on"
COPY docker/go/go.mod /go.mod
COPY docker/go/go.sum /go.sum
RUN go mod download

RUN #mkdir /app
WORKDIR /app

CMD go run cmd/app/main.go