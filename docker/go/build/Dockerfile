FROM golang:1.17 as dev

WORKDIR /app

COPY ./docker/go/go.mod ./
COPY ./docker/go/go.sum ./

RUN go mod download

CMD ["go", "run", "cmd/main.go"]

FROM golang:1.17 as builder

WORKDIR /app

COPY ./docker/go/go.mod ./
COPY ./docker/go/go.sum ./

RUN go mod download

COPY ./docker/go/  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /app/cmd/

RUN go build \
    -o /bin/main \
    -ldflags '-s -w'

FROM scratch as prod

WORKDIR /app/cmd

COPY --from=builder /bin/main /bin/main

ENTRYPOINT ["/bin/main"]