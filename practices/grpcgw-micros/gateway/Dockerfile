FROM golang:1.23-alpine AS builder
WORKDIR /app

# 复制代码和依赖
COPY gateway/main.go .
COPY product/gen product/gen
COPY third_party third_party
#COPY go.mod .  #If go.mod is there.

# 初始化模块并构建
RUN go mod init gateway && \
    go get github.com/grpc-ecosystem/grpc-gateway/v2  &&  \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o gateway .

FROM alpine:latest
COPY --from=builder /app/gateway .
EXPOSE 50050
CMD ["./gateway"]