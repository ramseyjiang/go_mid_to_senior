First Part:
0. Create a grpc folder, such as "greet".
1. create a grpcName which suffix with pb, such as "proto";
2. After that, create a file grpcName.proto, such as "greet.proto", which is in the proto folder.
3. Copy all code from greet.proto. Then, "cd ..", go back to greet folder.
4. run command following:
   % protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   proto/greet.proto

5. After you did 4, it will generate greet.pb.go file and greet_grpc.pb.go file automatically. These two files should not
   be edited manually. They only can be edit when you use the command in step 4.

Notice, "--go-grpc_opt=require_unimplemented_servers=false" is necessary! If not, it will have another method which not
implement in greet.proto. That method named "mustEmbedUnimplementedGreetServiceServer"

Second Part:
0. Create a folder name greet_server, and create a file named "server.go" in it;
1. Copy all server.go code in my github, you also can change some code if you understand how grpc works.
2. Create a folder name greet_client, and create a file named "client.go" in it;
3. Copy all client.go code in my github, you also can change some code if you understand how grpc works.

Third Part:(grpcui)
1. go get google.golang.org/grpc/reflection

2. import it in the server.go

3. Add below code in the main func. Please check 80-81 rows in the server.go
// Register reflection service on gRPC server.
reflection.Register(s)

4. Make server run, in greet folder
% go run server/server.go

5. After the above 3 steps, please run the following: (50051 is a portNumber.)
% grpcui -plaintext 0.0.0.0:50051
After that, you will show a grpcui webpage open automatically.

7. Make client run, in greet folder
% go run client/client.go

After all these above, you will see results in your terminal and in your grpcui webpage.