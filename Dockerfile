FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o opcua-mqtt-converter cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/opcua-mqtt-converter .

CMD ["./opcua-mqtt-converter"]