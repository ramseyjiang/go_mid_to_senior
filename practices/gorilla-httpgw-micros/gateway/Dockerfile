# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init gateway && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o gateway .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/gateway .
EXPOSE 8083
CMD ["./gateway"]
