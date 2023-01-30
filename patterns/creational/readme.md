Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing
code.

Creational Patterns includes Singleton, Builder, Factory, Prototype, and Abstract Factory Design Patterns.

The Singleton pattern is one of the most used design patterns out there or one of the easiest to grasp. It will provide
you with a single instance of an object, and guarantee that there are no duplicates. At first, call to use the instance,
it's created and then reused between all the parts in the application that need to use the particular behavior.

The Builder design pattern will reuse an algorithm to create many implementations of an interface. It means the builder
pattern will allow you to produce different types and representations of an object using the same construction code.

The Builder design pattern tries to:

1. Abstract complex creations so that object creation is separated from the object user.
2. Create an object step by step by filling its fields and creating the embedded objects.
3. Reuse the object creation algorithm between many objects.
