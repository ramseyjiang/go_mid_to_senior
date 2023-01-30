Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing
code.

Creational Patterns includes Singleton, Builder, Factory, Prototype, and Abstract Factory Design Patterns.

The Singleton pattern is one of the most used design patterns out there or one of the easiest to grasp. It will provide
a single instance of an object, and guarantee that there are no duplicates. At first, call to use the instance, if it's
created, and then reused between all the parts in the application that need to use the particular behavior.

The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic
they require. It will reuse an algorithm to create many implementations of an interface. It means the builder pattern
will allow you to produce different types and representations of an object using the same construction code. The builder
pattern's instance will provide the opening and closing braces {} and leaving the instance with zero values, or as
complex as an object that needs to make some API calls, check states, and create objects for its fields.

The Factory pattern is called Factory method pattern also. It is probably the second-best known and used design pattern
in the industry. The Factory pattern provides an interface for creating objects, but allows subclasses to decide which
class to instantiate. The interface also eases the process of downgrading or upgrading of the implementation of the
underlying type if needed.
