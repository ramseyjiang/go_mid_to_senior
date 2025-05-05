FROM golang:1.23-alpine  AS builder
WORKDIR /app

COPY product/proto product/proto
COPY product/main.go .
COPY third_party third_party
#COPY go.mod .  #If go.mod is there.

# Inital create model
RUN go mod init product-service &&  \
    go mod tidy &&  \
    CGO_ENABLED=0 GOOS=linux go build -o product-service .

FROM alpine:latest
COPY --from=builder /app/product-service .
EXPOSE 50051
CMD ["./product-service"]