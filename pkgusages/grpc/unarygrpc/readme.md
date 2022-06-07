What’s a Unary API?
a. Unary RPC calls are the basic Request / Response that everyone is familiar with
b. The client will send a message to the server and will receive one response from the server
c. Unary RPC calls will be the most common for your APIs.
d. Unary calls are very well suited when your data is small
e. Start with Unary when writing APIs and use streaming API if performance is an issue
f. In gRPC Unary Calls are defined using Protocol Buffers
g. For each RPC call, we have to define a Request message and a Response message


First Part:
0. Create a grpc folder, such as "unarygrpc".
1. create a grpcName which suffix with pb, such as "greetpb";
2. After that, create a file grpcName.proto, such as "greet.proto", which is in the greetpb folder.
3. Copy all code from greet.proto. Then, "cd ..", go back to unarygrpc folder.
4. run command following:
   % protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   greetpb/greet.proto

5. After you did 4, it will generate greet.pb.go file and greet_grpc.pb.go file automatically. These two files should not
   be edited manually. They only can be edit when you use the command in step 4.

Notice, "--go-grpc_opt=require_unimplemented_servers=false" is necessary! If not, it will have another method which not
implement in greet.proto. That method named "mustEmbedUnimplementedGreetServiceServer"

Second Part:
0. Create a folder name greet_server, and create a file named "server.go" in it;
1. Copy all server.go code in my github, you also can change some code if you understand how grpc works.
2. Create a folder name greet_client, and create a file named "client.go" in it;
3. Copy all client.go code in my github, you also can change some code if you understand how grpc works.

Third Part:
1. Make server run.
   In a terminal window, and in the greetUnaryGrpc folder, "% go run greet_server/server.go "
2. Make client run.
   In a terminal window, and in the greetUnaryGrpc folder, "% go run greet_client/client.go"