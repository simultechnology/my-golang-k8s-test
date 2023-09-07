FROM golang:1.21-alpine AS builder

WORKDIR /api

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY hello.go .

RUN go build -o app

FROM alpine

COPY --from=builder /api/app /app

CMD ["/app"]