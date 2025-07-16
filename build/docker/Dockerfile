FROM golang:1.24.3-alpine3.22 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/api/broker-service ./cmd/api

FROM alpine:3.22 AS app
RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
COPY --from=builder /app/bin/api/broker-service /usr/local/bin/
WORKDIR /usr/local/bin/
CMD ["./broker-service"]
