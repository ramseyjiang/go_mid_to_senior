1. for grpc gateway, you should get and install these following first:

go get -u google.golang.org/grpc go get -u google.golang.org/protobuf brew install protobuf-c

2. For google import ""google/api/annotations.proto", you should copy or download it first, then put it in your own
   folder, after that, you can import it. Otherwise, it cannot be found.

3. If your IDE is Goland, you should go to Preferences -> Language & Frameworks -> Protocol buffers, to set your own
   proto path. If you don't do that, your Goland won't find the proto path defined by yourself.
