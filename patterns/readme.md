Patterns list:

1. Dependency Injection Pattern
2. Fan In Out Pattern
3. Generator Pattern
4. Retry Pattern
5. Command Pattern
6. Semaphore Pattern

Creation patterns These patterns provide various object creation mechanisms, which increase flexibility and reuse of
existing code.

1. Singleton Pattern
2. Factory Pattern

Singleton Pattern is a class / type has only one instance and provides a global access point to it. Ensures that there
is a single instance of an object. This is useful for accessing a shared resource, such as a database.

Factory pattern is a creation design pattern that provides an interface for creating objects in a superclass, but allows
subclasses to alter the type of objects that will be created.

Builder pattern is a design pattern that provides a flexible solution to various object creation problems in
object-oriented programming. When you need to build multiple representations of a same object, or use the constructor
with a long parameter list or multiple constructors with different parameters, you should consider using it.

Struct patterns These patterns explain how to assemble objects and classes into larger structures while keeping these
structures flexible and efficient.

Behaviour patterns These patterns are concerned with algorithms and the assignment of responsibilities between objects.

Dependency Injection (DI) pattern is a design pattern used to implement IoC. It allows the creation of dependent objects
outside a class and provides those objects to a class through different ways.

Fan-in Fan-out is a way of Multiplexing and Demultiplexing in golang. Fan-in is a multiplexing strategy where the inputs
of several channels are combined to produce an output channel. Fan-out is demultiplexing strategy where a single channel
is split into multiple channels.

Generator Pattern is used to generate a sequence of values which is used to produce some output. This pattern is widely
used to introduce parallelism into loops. This allows the consumer of the data produced by the generator to run in
parallel when the generator function is busy computing the next value.

Retry Pattern is always used to make gracefully retry.
The pattern accounts for possible transient faults in a distributed system.
The Retry pattern, some form of backoff algorithm is implemented that increases the delay between each retry.

Command pattern is all about encapsulation and abstraction. 
In the case of the remote controller, each button is ideally encapsulated in that each button should be able to work independently of each other.
In terms of abstraction, this is the whole point that you don’t need to know the details of how exactly any button in the controller. 
You just want each button to be able to easily perform a task or action without having to think about the details.

Semaphore pattern is a variable or abstract data type used to control access to a common resource by multiple processes
in a concurrent system such as a multitasking operating system. In Golang, it can be implemented easily by buffered
channel. When buffered channel is full, the channel will lock the Goroutine and make it wait until a buffer becomes
available.

TODO:

1. Make routes which can access faninout.EntryInstance1() and faninout.EntryInstance2 directly.
2. Make a route which can access factory.Entry() directly.
3. Make a route which can access generator.Entry() directly.
4. Make a route which can access retry.Entry() directly.

CQRS is a Microservice architecture pattern, it stands for Command and Query Responsibility Segregation. The basic idea
behind this pattern is to keep write operations separate from the read operations. Instead of using one datastore to
perform CRUD operations, the read operation is performed on one datastore, and create/update is performed on a different
datastore. CQRS promises stability and scalability to large-scale applications along with significant performance
improvement.

Where is CQRS Used? a. CQRS is generally preferred in large-scale distributed applications that are ready-heavy. That is
if the number of reads in the application far outnumbers the number of write operations. b. When the read operations are
heavy and the datastore replicas could be placed near the geolocation where the application receives high traffic thus
improving the performance of database read operations.

How is CQRS implemented? Usually, a relational database is preferred for write operations, where all the data
constraints can be applied, a NoSQL database is used to support read operations. This means, that two different data
models are used for write and read operations.

Event-driven architecture Consider a scenario where a User calls a POST API, after that api call processes data are
updated into our relational database. Once the data is inserted into the DB a trigger is invoked that would update all
the event handlers to update the read database. When another User calls a GET API, the application would retrieve the
related data from the NoSQL database and send it back to the User.

The processing way as the image link: https://miro.medium.com/max/1400/1*lVt236fC4mL2Z7Wq8BZGaA.png
