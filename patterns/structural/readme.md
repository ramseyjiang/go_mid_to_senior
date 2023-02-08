Structural patterns, as the name implies, help us to shape our applications with commonly used structures and relationships. Structural design patterns
explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.

Composite pattern is a structural design pattern that allows composing objects into a tree-like structure and work with it as if it was a singular
object. It that lets you compose objects into tree structures and then work with these structures as if they were individual objects. In this pattern,
you can treat all objects as the same via the common interface.

Adapter pattern which allows incompatible objects/interfaces to collaborate. It is a special object that converts the interface of one object so that
another object can understand it. Adapters can not only convert data into various formats but can also help objects with different interfaces
collaborate.

Bridge pattern that lets you divides business logic or huge class into separate class hierarchies that can be developed independently. The Bridge
pattern lets you replace the implementation object inside the abstraction. The Bridge pattern tries to decouple things as usual with design patterns. It
decouples abstraction an object from its implementation the thing that the object does.

Proxy is a structural design pattern that provides an object that acts as a substitute or placeholder for a real service object used by a client. A
proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object. That means a proxy
controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.

Decorator is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside special wrapper objects, called
decorators. It allows you to decorate an already existing type with more functional features without actually touching it. When you think about
extending legacy code without the risk of breaking something, you should think of the Decorator pattern first.

Facade is a structural design pattern that provides a simplified(but limited) interface to a library, a framework, or any other complex set of classes
which contains lots of moving parts. A facade might provide limited functionality in comparison to working with the subsystem directly However, it
includes only those features that clients really care about.

Flyweight is a structural design pattern that allows programs to support huge quantities of objects by keeping their memory consumption low. The
Flyweight pattern takes out the common parts and creates flyweight objects.
