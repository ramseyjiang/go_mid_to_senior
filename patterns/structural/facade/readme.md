Facade is a structural design pattern that provides a simplified(but limited) interface to a library, a framework, or any other complex set of classes
which contains lots of moving parts. A facade might provide limited functionality in comparison to working with the subsystem directly. However, it
includes only those features that clients really care about.

Facade scenarios

1. Use the Facade pattern when you need to have a limited but straightforward interface to a complex subsystem.
2. Use the Facade when you want to structure a subsystem into layers. For example, Check account, Check security PIN, Credit/debit balance, Send
   notification, and so on.

How to Implement

1. Check whether it’s possible to provide a simpler interface than what an existing subsystem already provides.
2. Declare and implement this interface in a new facade class.
3. To get the full benefit from the pattern, make all the client code communicate with the subsystem only via the facade.
4. If the facade becomes too big, consider extracting part of its behavior to a new, refined facade class.

Pros: You can isolate your code from the complexity of a subsystem.

Cons: A facade can become a god object coupled to all classes of an app.

Facade is similar to Proxy in that both buffer a complex entity and initialize it on its own. Unlike Facade, Proxy has the same interface as its service
object, which makes them interchangeable.
