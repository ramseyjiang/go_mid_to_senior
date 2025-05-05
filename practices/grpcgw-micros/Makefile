.PHONY: generate

generate:
    # 安装protoc-gen-go和protoc-gen-grpc-gateway
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

	# Generate product service rpc code
	protoc	-I=product/proto \
		  	-I=third_party \
		  	--go_out=. \
		  	--go-grpc_out=. \
		  	--grpc-gateway_out=logtostderr=true:. \
		  	product/proto/product.proto

  	# Generate order service rpc code, it invokes from product proto
	protoc 	-I=. \
			-I=order/proto \
			-I=product/proto	\
			-I=third_party	\
			--go_out=.	\
			--go-grpc_out=.	\
			--grpc-gateway_out=logtostderr=true:.	\
			order/proto/order.proto

	# format generation code
	gofmt -w product/gen/*.go order/gen/*.go