Decorator is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside special wrapper objects, called
decorators. It allows you to decorate an already existing type with more functional features without actually touching it. When you think about
extending legacy code without the risk of breaking something, you should think of the Decorator pattern first.

“Wrapper” is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern. A wrapper is an object that can be
linked with some targets object. The wrapper contains the same set of methods as the target and delegates to it all requests it receives. However, the
wrapper may alter the result by doing something either before or after it passes the request to the target.

The Decorator type implements the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack
as many decorators (dolls) as you want by simply storing the old decorator in a field of the new one.

How to Implement

1. Make sure your business domain can be represented as a primary component with multiple optional layers over it.
2. Figure out what methods are common to both the primary component and the optional layers. Create a component interface and declare those methods
   there.
3. Create a concrete component class and define the base behavior in it.
4. Create a base decorator class. It should have a field for storing a reference to a wrapped object. The field should be declared with the component
   interface type to allow linking to concrete components as well as decorators. The base decorator must delegate all work to the wrapped object.
5. Make sure all classes implement the component interface.
6. Create concrete decorators by extending them from the base decorator. A concrete decorator must execute its behavior before or after the call to the
   parent method (which always delegates to the wrapped object).
7. The client code must be responsible for creating decorators and composing them in the way the client needs.

Pros

1. You can extend an object’s behavior without making a new subclass.
2. You can add or remove responsibilities from an object at runtime.
3. You can combine several behaviors by wrapping an object into multiple decorators.
4. Single Responsibility Principle. You can divide a monolithic class that implements many possible variants of behavior into several smaller classes.

Cons

1. It’s hard to remove a specific wrapper from the wrappers stack.
2. It’s hard to implement a decorator in such a way that its behavior doesn’t depend on the order in the decorators stack.
3. The initial configuration code of layers might look pretty ugly.

Difference between decorator pattern and proxy pattern. The Decorator pattern is used to add additional responsibilities to an object dynamically. It
provides a flexible alternative to subclassing for extending functionality. The decorator objects are wrapped around the original object and the
original object is not modified. The decorator pattern is often used for implementing cross-cutting concerns, such as logging or security, in a clean
and reusable way.

The Proxy pattern, on the other hand, provides a placeholder for another object to control access to it. The proxy object acts as an intermediary
between the client and the real subject. This pattern is often used for creating objects on demand, for remote communication, or for performance
optimization.

In summary, the Decorator pattern is focused on adding functionality to an object, while the Proxy pattern is focused on controlling access to an
object.
