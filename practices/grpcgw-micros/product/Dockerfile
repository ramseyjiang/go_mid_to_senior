FROM golang:1.23-alpine  AS builder
WORKDIR /app

COPY product/proto product/proto
COPY product/main.go .
COPY third_party third_party
#COPY go.mod .  #If go.mod is there.

# 初始化模块并构建
RUN go mod init product-service &&  \
    go get google.golang.org/grpc &&  \
    CGO_ENABLED=0 GOOS=linux go build -o product-service .

FROM alpine:latest
COPY --from=builder /app/product-service .
EXPOSE 50051
CMD ["./product-service"]