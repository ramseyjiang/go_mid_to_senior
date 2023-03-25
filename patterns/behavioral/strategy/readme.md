Strategy is a behavioral design pattern that allows you to define a family of algorithms and dynamically switch between them at runtime. Strategy lets
the algorithm vary independently of clients that use it. The Strategy pattern is probably the easiest to understand of the Behavioral patterns.

Objectives

1. Define a common interface for multiple algorithms or behaviors.
2. Encapsulate each algorithm or behavior in a separate class.
3. Allow the client code to select the desired algorithm or behavior at runtime.

Pros:

1. Better separation of concerns: Each strategy encapsulates its own logic, making the code more modular and easier to maintain.
2. Flexibility: New strategies can be added or existing ones modified without affecting the client code.
3. Easy to switch between algorithms at runtime.

Cons:

1. Increase the number of classes and make the code more complex
2. Make the code harder to understand if too many strategies are used

How to implement

1. Define a strategy interface that represents the algorithm you want to encapsulate.
2. Implement concrete strategies that adhere to the above interface.
3. Create a context class that uses the Strategy Interface:
4. Use the context class and concrete strategies in client code
