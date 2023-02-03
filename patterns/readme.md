Design patterns are typical solutions to commonly occurring problems in software design. They are like pre-made blueprints that you can
customize to solve a recurring design problem in your code.

Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code. Creational
Patterns includes Singleton, Builder, Factory, Prototype, and Abstract Factory Design Patterns.

The Singleton pattern is one of the most used design patterns out there or one of the easiest to grasp. It will provide a single instance of
an object, and guarantee that there are no duplicates. At first, call to use the instance, if it's created, and then reused between all the
parts in the application that need to use the particular behavior.

The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic they require. It
will reuse an algorithm to create many implementations of an interface. It means the builder pattern will allow you to produce different
types and representations of an object using the same construction code. The builder pattern's instance will provide the opening and closing
braces {} and leaving the instance with zero values, or as complex as an object that needs to make some API calls, check states, and create
objects for its fields.

The Factory pattern is called Factory method pattern also. It is probably the second-best known and used design pattern in the industry. The
Factory pattern provides an interface for creating objects, but allows subclasses to decide which class to instantiate. The interface also
eases the process of downgrading or upgrading of the implementation of the underlying type if needed.

The Abstract pattern is called Abstract Factory pattern also. Abstract pattern is used to solve the problem of creating entire product
families without specifying their concrete classes. It provides an interface for creating all distinct products but leaves the actual product
creation to concrete factory classes. Each factory type corresponds to a certain product variety.

The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface
makes the client code independent of the concrete classes of objects that it clones. The aim of the Prototype pattern is to have an object or
a set of objects that is already created at compilation time, but you can clone as many times as you want at runtime.An object that supports
cloning is called a prototype. Here’s how it works: you create a set of objects, configured in various ways. When you need an object like the
one you’ve configured, you just clone a prototype instead of constructing a new object from scratch.

It's worth mentioning that these patterns are not thread-free.

Patterns list:

1. Dependency Injection Pattern
2. Fan In Out Pattern
3. Generator Pattern
4. Retry Pattern
5. Command Pattern
6. Semaphore Pattern

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
