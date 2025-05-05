FROM golang:1.23 AS builder
WORKDIR /app

# Copy Code and dependencies
COPY order/proto order/proto
COPY product/proto product/proto
COPY third_party third_party
COPY order/main.go .
#COPY go.mod .  #If go.mod is there.

# Inital modular and create
RUN go mod init order-service && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o order-service .

FROM alpine:latest
COPY --from=builder /app/order-service .
EXPOSE 50052
CMD ["./order-service"]