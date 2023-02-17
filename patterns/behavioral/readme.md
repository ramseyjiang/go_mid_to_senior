Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects.

Strategy is a behavioral design pattern that allows you to define a family of algorithms and dynamically switch between them at runtime. Strategy lets
the algorithm vary independently of clients that use it. The Strategy pattern is probably the easiest to understand of the Behavioral patterns.

Chain of Responsibility is a behavioral design pattern that allows multiple objects to handle a request, with each object having the opportunity to
either handle the request or pass it along to the next object in the chain. The objects in the chain are not aware of each other, but instead
communicate through a common interface, allowing new objects to be added or removed from the chain easily.

Command is a behavioral design pattern that allows encapsulating a request or operation as an object. It converts requests or simple operations into a
stand-alone object that can be stored, passed, and executed independently of the original requester. The pattern separates the requester (client) from
the object that performs the action (receiver). This transformation lets you pass requests as a method arguments, delay or queue a request’s execution,
and support undoable operations.

The Template pattern is a behavioral design pattern that defines the skeleton of an algorithm in a base class and allows subclasses to override some
steps of the algorithm without changing its structure. In other words, it provides a way to define the steps of an algorithm, while allowing subclasses
to implement specific behaviors that can vary.

The Memento pattern is a behavioral design pattern that allows an object to save its internal state so that it can be restored later, without violating
encapsulation. In other words, it provides a way to capture the current state of an object and restore it to that state later on.

The Iterator pattern is a behavioral design pattern that provides a way to traverse a collection of objects without exposing the underlying
implementation details. It allows you to access the elements of a collection sequentially, one at a time, without having to know the internal structure
of the collection.

The Interpreter pattern is widely used to solve business cases where it's useful to have a language to perform common operations. The pattern involves
creating a set of algorithm classes that are used to solve logical puzzles.

Visitor is a behavioral design pattern that allows you to add new algorithms from the objects structure on which they operate. It involves defining a
separate object (the visitor) that can visit each object in a complex structure and perform some operation on it. This way, you can modify the behavior
of a group of objects without changing their individual classes.
