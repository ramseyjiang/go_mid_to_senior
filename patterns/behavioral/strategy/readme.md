Strategy is a behavioral design pattern that allows you to define a family of algorithms and dynamically switch between them at runtime. Strategy lets
the algorithm vary independently of clients that use it. The Strategy pattern is probably the easiest to understand of the Behavioral patterns.

Objectives

1. Encapsulate algorithms in separate objects to allow for dynamic and interchangeable behavior
2. All types achieve the same functionality in a different way but the client of the strategy isn't affected.

Pros:

1. Promotes clean, modular, and reusable code by encapsulating algorithms in separate objects
2. Increases flexibility by allowing the algorithm to be changed at runtime
3. Enables better separation of concerns by isolating algorithms from the context in which they are used

Cons:

1. Increase the number of classes and make the code more complex
2. Make the code harder to understand if too many strategies are used

How to implement

1. Define an interface that represents the algorithm you want to encapsulate.
2. Define the methods that are needed to perform the algorithm in the above interface.
3. Create separate structs that implement the interface, each defining a specific algorithm.
4. The context class would then have a reference to an instance of the interface, and use it to perform the algorithm.
