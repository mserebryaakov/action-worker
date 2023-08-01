FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/main cmd/main.go

FROM alpine:3.18

COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]