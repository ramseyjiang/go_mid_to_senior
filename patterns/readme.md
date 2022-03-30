Patterns list:
1. Dependency Injection Pattern
2. Factory Pattern
3. Fan In Out Pattern
4. Generator Pattern
5. Retry Pattern
6. Singleton Pattern
7. Command Pattern
8. Semaphore Pattern

Dependency Injection (DI) pattern is a design pattern used to implement IoC. 
It allows the creation of dependent objects outside a class and provides those objects to a class through different ways.

Factory pattern is a creation design pattern that provides an interface for creating objects in a superclass,
but allows subclasses to alter the type of objects that will be created.

Fan-in Fan-out is a way of Multiplexing and Demultiplexing in golang.
Fan-in is a multiplexing strategy where the inputs of several channels are combined to produce an output channel. 
Fan-out is demultiplexing strategy where a single channel is split into multiple channels.

Generator Pattern is used to generate a sequence of values which is used to produce some output.
This pattern is widely used to introduce parallelism into loops.
This allows the consumer of the data produced by the generator to run in parallel when the generator function is busy computing the next value.

Retry Pattern is always used to make gracefully retry.
The pattern accounts for possible transient faults in a distributed system.
The Retry pattern, some form of backoff algorithm is implemented that increases the delay between each retry.

Singleton Pattern is a class / type has only one instance and provides a global access point to it.
Ensures that there is a single instance of an object. This is useful for accessing a shared resource, such as a database.

Command pattern is all about encapsulation and abstraction. 
In the case of the remote controller, each button is ideally encapsulated in that each button should be able to work independently of each other.
In terms of abstraction, this is the whole point that you don’t need to know the details of how exactly any button in the controller. 
You just want each button to be able to easily perform a task or action without having to think about the details.

Semaphore pattern is a variable or abstract data type used to control access to a common resource by multiple processes 
in a concurrent system such as a multitasking operating system.
In Golang, it can be implemented easily by buffered channel.
When buffered channel is full, the channel will lock the Goroutine and make it wait until a buffer becomes available.

TODO:
1. Make routes which can access faninout.EntryInstance1() and faninout.EntryInstance2 directly.
2. Make a route which can access factory.Entry() directly.
3. Make a route which can access generator.Entry() directly.
4. Make a route which can access retry.Entry() directly.