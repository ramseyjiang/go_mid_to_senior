What is gRPC?
gRPC is a modern open-source high-performance Remote Procedure Call (RPC) framework that can run in any environment.
It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing,
health checking, and authentication. It is also applicable in the last mile of distributed computing to connect devices,
mobile applications, and browsers to backend services.

There are 4 types of gRPC: unary, server-streaming, client-streaming, and bidirectional streaming.

1. Unary: A method is unary type where the client sends a single request to the server and gets a single response back,
just like a normal function call. A unary call is complete when the response is returned.
Unary RPC calls are well suited, when data to be transferred is small. 
It is available even with Http1 and Http2. Others are available with Http2 only.
Its format is "rpc SayHello(HelloRequest) returns (HelloResponse)".
Unary is what a tradition API looks like HTTP REST.
HTTP2 as we have seen, enable APIs to have streaming capabilities.

2. Server streaming: A server streaming method is where the client sends a request to the server and gets a stream to read a sequence
of messages back. A server streaming call is complete when the method returns.
Its format is "rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse)".

3. Client streaming: A Client streaming method is where the client writes a sequence of messages and sends them to the server.
Once the client has finished writing messages, it waits for the server to read them and return its response.
Its format is "rpc LotsOfRequests(stream HelloRequest) returns (HelloResponse)".

4. Bi-directional streaming: A Bidirectional streaming where both sides send a sequence of messages using a read-write stream.
The two streams operate independently, so clients and servers can read and write in whatever order they like.
Its format is "rpc BidiHello(stream HelloRequest) returns (stream HelloResponse)".


gRPC Servers are async by default. 
This means they don't block threads on requests.
Therefore, each gRPC server can serve millions of requests in parallel.

gRPC Clients can be async or sync(blocking)
Clients can decide which model works best for the performance needs.
gRPC Clients can perform client side load balancing.

For security, gRPC strongly recommend for you to use SSL(encryption the wire) in your API.


The differences between REST & gRPC :

1. REST shall be using JSON as data-transfer standard which would be slower, bigger and text-based. W
While gRPC uses the protobuf as data-transfer standard which are smaller and faster in nature.

2. REST uses the HTTP/1.1 while gRPC uses the HTTP/2 which is lot much faster and efficient.

3. REST only supports Client to server calls while gRPC supports bidirectional and async calls as well.

4. REST only supports request/response while gRPC supports streaming capabilities as well.

5. REST is purely Resource-oriented while gRPC is purely API oriented with free-design.

6. REST supports auto-code-generation using swagger and openAPI as 2nd class citizens.
While gRPC supports auto-code-generation using protobuf as 1st class citizens.

7. REST is Verbs based, and thus we have to basic plumbing ourselves.
While gRPC is RPC based i.e. we can invoke the functions at server easily.


Five explicit GRPC good points.
1. Easy code definition in over languages.
2. Use a modern, low latency HTTP/2 transport mechanism
3. SSL Security is built in
4. Support for streaming APIs for maximum performance
5. GRPC is API oriented, instead of Resource Oriented like REST.


What’s a Unary API?
a. Unary RPC calls are the basic Request / Response that everyone is familiar with
b. The client will send a message to the server and will receive one response from the server
c. Unary RPC calls will be the most common for your APIs.
d. Unary calls are very well suited when your data is small
e. Start with Unary when writing APIs and use streaming API if performance is an issue
f. In gRPC Unary Calls are defined using Protocol Buffers
g. For each RPC call, we have to define a Request message and a Response message


Size: JSON will produce an object that is larger than binary object.
JSON supports only a limited number of data types:string, number, boolean, null, object, array.
Binary supports more.