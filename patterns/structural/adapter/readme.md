The Adapter pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It converts the interface of one
class into another interface that the client expects. This allows objects with incompatible interfaces to work together, without modifying their source
code. It means the adaptor pattern uses a special algorithm to convert one object's interface so that another object can understand it. Adapters can not
only convert data into various formats but can also help objects with different interfaces collaborate.

Objectives

1. Allow objects with incompatible interfaces to collaborate.
2. Provide a bridge between two incompatible interfaces, allowing them to work together.
3. Adapt an existing interface to meet the requirements of a new interface.

The adapter pattern comprises the target, adaptee, adapter, and client.

1. The Target is the interface that the client wants to call and invoke.
2. The Adaptee is the class that needs to be adapted.
3. The Adapter class implements the above target interface.
4. The Client orchestrates the adapter by calling the adaptee’s method indirectly.

How to Implement

1. Make sure that you have at least two classes with incompatible interface
2. Declare the client interface and describe how clients communicate with the service.
3. Create the adapter class and make it follow the client interface. Leave all the methods empty for now.
4. Add a field to the adapter class to store a reference to the service object. The common practice is to initialize this field via the constructor, but
   sometimes it’s more convenient to pass it to the adapter when calling its methods.
5. One by one, implement all methods of the client interface in the adapter class. The adapter should delegate most of the real work to the service
   object, handling only the interface or data format conversion.
6. Clients should use the adapter via the client interface. This will let you change or extend the adapters without affecting the client code.

Pros

1. Single Responsibility Principle. You can separate the interface or data conversion code from the primary business logic of the program.
2. Open/Closed Principle. You can introduce new types of adapters into the program without breaking the existing client code, as long as they work with
   the adapters through the client interface.

Cons

The overall complexity of the code increases because you need to introduce a set of new interfaces and classes. Sometimes it’s simpler just to change
the service class so that it matches the rest of your code.

func main() { http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello, World")
})
log.Fatal(http.ListenAndServe(":8080", nil))
} The HandleFunc function is actually part of an adapter for using functions directly as ServeHTTP implementations.
