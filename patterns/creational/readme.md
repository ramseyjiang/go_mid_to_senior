Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing
code. Creational Patterns includes Singleton, Builder, Factory, Prototype, and Abstract Factory Design Patterns.

The Singleton pattern is one of the most used design patterns out there or one of the easiest to grasp. It will provide
a single instance of an object, and guarantee that there are no duplicates. At first, call to use the instance, if it's
created, and then reused between all the parts in the application that need to use the particular behavior.

The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic
they require. It will reuse an algorithm to create many implementations of an interface. It means the builder pattern
will allow you to produce different types and representations of an object using the same construction code. The builder
pattern's instance will provide the opening and closing braces {} and leaving the instance with zero values, or as
complex as an object that needs to make some API calls, check states, and create objects for its fields.

The Factory pattern is called Factory method pattern also. It is probably the second-best known and used design pattern in the industry. The
Factory pattern provides an interface for creating objects, but allows subclasses to decide which class to instantiate. The interface also
eases the process of downgrading or upgrading of the implementation of the underlying type if needed.

The Abstract pattern is called Abstract Factory pattern also. Abstract pattern is used to solve the problem of creating entire product families without
specifying their concrete classes. It provides an interface for creating all distinct products but leaves the actual product creation to concrete
factory classes. Each factory type corresponds to a certain product variety.

The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface makes the
client code independent of the concrete classes of objects that it clones. The aim of the Prototype pattern is to have an object or a set of objects
that is already created at compilation time, but you can clone as many times as you want at runtime.An object that supports cloning is called a
prototype. Here’s how it works: you create a set of objects, configured in various ways. When you need an object like the one you’ve configured, you
just clone a prototype instead of constructing a new object from scratch.

The Multiton pattern is a creational design pattern that is similar to the Singleton pattern, but with a twist. The Singleton pattern ensures that a
class has only one instance, while the Multiton pattern ensures that a class has a limited number of instances, each with its own unique identifier. The
Multiton pattern is useful in situations where multiple instances of a class are required, but each instance must have a unique identity.

It's worth mentioning that these patterns are not thread-free.

----------------------------------------------------------------------------------------------------------------------------------
we have seen the differences when approaching the same problem with two different solutions–the Abstract factory and the Builder pattern. As you have
seen, with the Builder pattern, we had an unstructured list of objects (cars with motorbikes in the same factory). Also, we encouraged reusing the
building algorithm in the Builder pattern. In the Abstract factory, we have a very structured list of vehicles (the factory for motorbikes and a factory
for cars). We also didn't mix the creation of cars with motorbikes, providing more flexibility in the creation process. The Abstract factory and Builder
patterns can both resolve the same problem, but your particular needs will help you find the slight differences that should lead you to take one
solution or the other.
