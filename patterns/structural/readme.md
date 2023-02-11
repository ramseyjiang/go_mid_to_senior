Structural design patterns are a category of design patterns that deal with object composition and class relationships and are used to simplify the
design of large and complex systems. These patterns concern how classes and objects can be combined to form larger structures and provide a way to
manage the relationships between entities to simplify the design of complex systems.

1. Composite pattern is a structural design pattern that allows composing objects into a tree-like structure and working with it as if it was a singular
   object. It lets you compose objects into tree structures and then work with these structures as if they were individual objects. In this pattern, you
   can treat all objects as the same via the common interface.
2. Adapter pattern which allows incompatible objects/interfaces to collaborate. It is a special object that converts the interface of one object so that
   another object can understand it. Adapters can not only convert data into various formats but can also help objects with different interfaces
   collaborate.
3. Bridge pattern that lets you divide business logic or huge class into separate class hierarchies that can be developed independently. The Bridge
   pattern lets you replace the implementation object inside the abstraction. The Bridge pattern tries to decouple things as usual with design patterns.
   It decouples the abstraction of an object from its implementation of the thing that the object does.
4. Proxy pattern that provides an object that acts as a substitute or placeholder for a real service object used by a client. A proxy receives client
   requests, does some work (access control, caching, etc.) and then passes the request to a service object. That means a proxy controls access to the
   original object, allowing you to perform something either before or after the request gets through to the original object.
5. Decorator pattern that allows adding new behaviours to objects dynamically by placing them inside special wrapper objects, called decorators. It
   allows you to decorate an already existing type with more functional features without actually touching it. When you think about extending legacy
   code without the risk of breaking something, you should think of the Decorator pattern first.
6. Facade pattern that provides a simplified(but limited) interface to a library, a framework, or any other complex set of classes that contains lots of
   moving parts. A facade might provide limited functionality in comparison to working with the subsystem directly However, it includes only those
   features that clients really care about.
7. Flyweight pattern that allows programs to support huge quantities of objects by keeping their memory consumption low. The Flyweight pattern takes out
   the common parts and creates flyweight objects.

In conclusion, Structural design patterns are used to simplify the design of complex systems by breaking them down into smaller, more manageable parts.
They provide a way to manage the relationships between objects and simplify the design of complex systems by making it easier to add, remove, or modify
parts of the system without having to make significant changes to the existing implementation.

These patterns are commonly used in software development to provide a flexible and modular architecture, which makes it easier to add new features or
functionality to a system. They provide a way to manage the relationships between objects in a system and simplify the design of complex systems by
breaking them down into smaller, more manageable parts.
