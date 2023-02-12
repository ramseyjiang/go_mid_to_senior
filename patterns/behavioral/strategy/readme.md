Strategy is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the
algorithm vary independently of clients that use it. The Strategy pattern is probably the easiest to understand of the Behavioral patterns.

Objectives

1. Provide a few algorithms to achieve some specific functionality.
2. All types achieve the same functionality in a different way but the client of the strategy isn't affected.

Pros:

1. The main advantage of using the Strategy pattern is that it allows for the separation of concerns. The logic for a specific behavior can be
   encapsulated within a separate strategy class, making it easier to maintain and understand.
2. Another benefit of using the Strategy pattern is that it makes it easy to change the behavior of a class at runtime. This can be useful in cases
   where the behavior of a class needs to be dynamic and can change based on user input or some other condition.
3. The Strategy pattern also allows for the reuse of code. By encapsulating the logic for a specific behavior in a separate class, that class can be
   used in multiple places throughout the application.

Cons:

1. One potential drawback of using the Strategy pattern is that it can result in an increase in the number of classes in the application. This can make
   it more difficult to understand the relationships between classes and how they interact with each other.
2. Another potential disadvantage of the Strategy pattern is that it can make the code more complex if not used correctly. It's important to carefully
   consider the design of the strategies and how they will interact with the context class.
