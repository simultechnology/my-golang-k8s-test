FROM golang:1.21-alpine

WORKDIR /api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go .

RUN go build -o app

CMD ["./app"]