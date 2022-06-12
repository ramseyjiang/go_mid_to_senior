This is a unary gRPC.
This way is very similar as REST.

Execute following command in the unarycal folder
% protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/*