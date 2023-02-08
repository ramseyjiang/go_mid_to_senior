Flyweight is a structural design pattern that allows programs to support huge quantities of objects by keeping their memory consumption low. The
Flyweight pattern takes out the common parts and creates flyweight objects.

Scenario

Use the Flyweight pattern only when your program must support a huge number of objects which barely fit into available RAM.

How to Implement

1. Divide fields of a class that will become a flyweight into two parts. The intrinsic state: the fields that contain unchanging data duplicated across
   many objects. The extrinsic state: the fields that contain contextual data unique to each object.
2. Leave the fields that represent the intrinsic state in the class, but make sure they’re immutable.
3. Go over methods that use fields of the extrinsic state. For each field used in the method, introduce a new parameter and use it instead of the field.
4. Optionally, create a factory class to manage the pool of flyweights. It should check for an existing flyweight before creating a new one.
5. The client must store or calculate values of the extrinsic state (context) to be able to call methods of flyweight objects. For the sake of
   convenience, the extrinsic state along with the flyweight-referencing field may be moved to a separate context class.

Pros

You can save lots of RAM, assuming your program has tons of similar objects.

Cons

1. You might be trading RAM over CPU cycles when some context data needs to be recalculated each time somebody calls a flyweight method.
2. The code becomes much more complicated. New team members will always be wondering why the state of an entity was separated in such a way.

Since the same flyweight object can be used in different contexts, you have to make sure that its state can’t be modified. A flyweight should initialize
its state just once, via constructor parameters. It shouldn’t expose any setters or public fields to other objects.

When you use the Flyweight pattern, it is very common to have a Flyweight factory, which uses other types of creational patterns to retrieve the objects
it needs.

Difference the Singleton pattern with the Flyweight pattern. With the Singleton pattern, we ensure that the same type is created only once. Also, the
Singleton pattern is a Creational pattern. With Flyweight, which is a Structural pattern, we aren't worried about how the objects are created, but about
how to structure a type to contain heavy information in a light way.
