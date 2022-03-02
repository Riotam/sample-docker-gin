FROM golang:1.17

COPY docker/go/app/org_src/go.mod /go.mod
COPY docker/go/app/org_src/go.sum /go.sum
WORKDIR /app
RUN ls
RUN go mod download

COPY ../.. /app

CMD go run org_src/main.go