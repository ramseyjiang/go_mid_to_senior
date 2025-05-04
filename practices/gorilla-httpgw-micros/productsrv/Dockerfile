# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init product-service && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o product-service .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/product-service .
EXPOSE 8082
CMD ["./product-service"]
