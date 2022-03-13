FROM golang:1.17 as dev

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]

FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN go build \
    -o /bin/main \
    -ldflags '-s -w'

FROM scratch as prod

WORKDIR /app

COPY --from=builder /bin/main /bin/main

ENTRYPOINT ["/bin/main"]