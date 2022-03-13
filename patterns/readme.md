Patterns list:
1. Dependency Injection Pattern
2. Factory Pattern
3. Fan In Out Pattern
4. Generator Pattern
5. Retry Pattern
6. Singleton Pattern

Singleton Pattern is a class / type has only one instance and provides a global access point to it.
Ensures that there is a single instance of an object. This is useful for accessing a shared resource, such as a database.

Factory pattern is a creational design pattern that provides an interface for creating objects in a superclass,
but allows subclasses to alter the type of objects that will be created.

Fan-in Fan-out is a way of Multiplexing and Demultiplexing in golang.
Fan-in refers to processing multiple input data and combining into a single entity.
Fan-out is dividing the data into multiple smaller chunks, distributing the work amongst a group of workers to parallelize CPU use and I/O.

Generator Pattern is used to generate a sequence of values which is used to produce some output.
This pattern is widely used to introduce parallelism into loops.
This allows the consumer of the data produced by the generator to run in parallel when the generator function is busy computing the next value.

Retry Pattern is always used to make gracefully retry.
The pattern accounts for possible transient faults in a distributed system.
The Retry pattern, some form of backoff algorithm is implemented that increases the delay between each retry.

TODO:
1. Make routes which can access faninout.EntryInstance1() and faninout.EntryInstance2 directly.
2. Make a route which can access factory.Entry() directly.
3. Make a route which can access generator.Entry() directly.
4. Make a route which can access retry.Entry() directly.