The Mediator pattern is a behavioral design pattern that allows communication between different objects through a central object called the Mediator.
This pattern reduces coupling between objects by making objects communicate indirectly. The Mediator object encapsulates the communication logic between
objects and acts as an intermediary to coordinate their interactions.

Objectives

1. To reduce the coupling between objects and promote better code organization by encapsulating the communication logic in a central object.
2. To simplify the communication between objects by providing a single interface that each object can use to interact with other objects.
3. To make it easier to add new objects to the system by eliminating the need to modify the interactions between existing objects.
4. To improve the maintainability of the codebase by making it easier to understand and modify the communication logic between objects.

Pros

1. Loose coupling between objects.
2. Improved maintainability and scalability of the codebase.
3. Simplified communication between objects.
4. Increased flexibility and extensibility.

Cons

1. Become complex if it needs to handle a large number of interactions between objects.
2. Become a performance bottleneck if it needs to process a large volume of requests.

How to implement

1. Define an interface that represents the mediator. This interface should include methods for registering and communicating with the objects that will
   be using the mediator.
2. Implement the mediator interface with a concrete type. This type should maintain a list of registered objects and be responsible for routing messages
   between them.
3. Define the objects that will be using the mediator. These objects should have a reference to the mediator and use it to communicate with each other.
4. Register the objects with the mediator. This can be done either in the constructor for each object, or through a separate registration method.
5. When an object needs to communicate with another object, it should send a message to the mediator, which will then route it to the appropriate
   recipient.

The Mediator pattern is useful when you have a system where objects need to communicate with each other in complex ways. By introducing a mediator, you
can simplify the communication between objects and decouple their relationships. Mediator pattern will act as the type in charge of exchanging
communication between two objects. This way, the communicating objects don't need to know each other and can change more freely.
