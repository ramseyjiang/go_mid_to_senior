.PHONY: generate

generate:
    # 安装protoc-gen-go和protoc-gen-grpc-gateway
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

	# 生成代码
	protoc -I=product/proto \
		   -I=third_party \
		   --go_out=. \
		   --go-grpc_out=. \
		   --grpc-gateway_out=logtostderr=true:. \
		   product/proto/product.proto

	# 格式化生成的代码
	gofmt -w product/gen/*.go