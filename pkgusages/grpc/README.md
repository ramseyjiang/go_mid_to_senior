What is gRPC?
gRPC is a modern open-source high-performance Remote Procedure Call (RPC) framework that can run in any environment.
It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing,
health checking, and authentication. It is also applicable in the last mile of distributed computing to connect devices,
mobile applications, and browsers to backend services.

There are 4 types of gRPC: unary, server-streaming, client-streaming, and bidirectional streaming.

1. Unary: A method is unary type where the client sends a single request to the server and gets a single response back,
just like a normal function call. A unary call is complete when the response is returned.
Its format is "rpc SayHello(HelloRequest) returns (HelloResponse)".

2. Server streaming: A server streaming method is where the client sends a request to the server and gets a stream to read a sequence
of messages back. A server streaming call is complete when the method returns.
Its format is "rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse)".

3. Client streaming: A Client streaming method is where the client writes a sequence of messages and sends them to the server.
Once the client has finished writing messages, it waits for the server to read them and return its response.
Its format is "rpc LotsOfRequests(stream HelloRequest) returns (HelloResponse)".

4. Bi-directional streaming: A Bidirectional streaming where both sides send a sequence of messages using a read-write stream.
The two streams operate independently, so clients and servers can read and write in whatever order they like.
Its format is "rpc BidiHello(stream HelloRequest) returns (stream HelloResponse)".


What’s a Unary API?
a. Unary RPC calls are the basic Request / Response that everyone is familiar with
b. The client will send a message to the server and will receive one response from the server
c. Unary RPC calls will be the most common for your APIs.
d. Unary calls are very well suited when your data is small
e. Start with Unary when writing APIs and use streaming API if performance is an issue
f. In gRPC Unary Calls are defined using Protocol Buffers
g. For each RPC call, we have to define a Request message and a Response message