The Factory pattern is called Factory method pattern also. It is probably the second-best known and used design pattern
in the industry. The Factory pattern provides an interface for creating objects, but allows subclasses to decide which
class to instantiate. The interface also eases the process of downgrading or upgrading of the implementation of the
underlying type if needed.

Objectives

1. Delegating the creation of new instances of structures to a different part of the program
2. Working at the interface level instead of with concrete implementations
3. Grouping families of objects to obtain a family object creator

Pros

1. You avoid tight coupling between the creator and the concrete products.
2. Single Responsibility Principle. You can move the product creation code into one place in the program, making the
   code easier to support.
3. Open/Closed Principle. You can introduce new types of products into the program without breaking existing client
   code.

Cons The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern.
