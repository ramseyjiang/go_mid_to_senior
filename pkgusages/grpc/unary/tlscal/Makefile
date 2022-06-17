.PHONY: cert
cert:
		cd cert; sh ./gen.sh; cd ..

.PHONY: server
server:
		go run server/server.go

.PHONY: client
client:
		go run client/client.go

.PHONY: generate
generate:
		protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*