Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects.

Strategy is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the
algorithm vary independently of clients that use it. The Strategy pattern is probably the easiest to understand of the Behavioral patterns.

Chain of Responsibility is a behavioral design pattern that allows multiple objects to handle a request, with each object having the opportunity to
either handle the request or pass it along to the next object in the chain. The objects in the chain are not aware of each other, but instead
communicate through a common interface, allowing new objects to be added or removed from the chain easily.

Command is a behavioral design pattern that allows encapsulating a request or operation as an object. It converts requests or simple operations into a
stand-alone object that can be stored, passed, and executed independently of the original requester. The pattern separates the requester (client) from
the object that performs the action (receiver). This transformation lets you pass requests as a method arguments, delay or queue a request’s execution,
and support undoable operations.
