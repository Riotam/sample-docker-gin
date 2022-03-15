FROM golang:1.17 as dev

WORKDIR /app

COPY . .

RUN go mod tidy

# ホットリロードモードで実行
RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["air"]

FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

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