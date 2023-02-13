The Chain of Responsibility pattern is a behavioral design pattern that provides a way to process a sequence of requests, where each request is passed
through a chain of objects until it is handled by one of them. The objects in the chain do not know each other, but instead, they communicate through a
common interface.

Objectives

1. To decouple the sender of a request from its receiver.
2. To allow for the dynamic modification of the chain of objects that handle requests.
3. To promote the single responsibility principle by allowing objects to handle only the requests they are responsible for.

Pros

1. Loose coupling: The objects in the chain are decoupled from each other, so changes in one object do not affect the others.
2. Increased flexibility: New objects can be added to the chain or existing objects can be removed, making it easy to change the behavior of the chain.
3. Improved maintainability: The logic of handling requests can be changed easily by changing the order of objects in the chain or adding/removing
   objects from the chain.

Cons

1. Performance overhead: If there are many objects in the chain, passing requests through them can become slow and resource-intensive.
2. Increased complexity: Maintaining the chain can become complex as the number of objects in the chain increases.
3. Debugging difficulties: Debugging the chain can be difficult because of its dynamic nature.

In summary, the Chain of Responsibility pattern is a powerful tool for processing requests in an object-oriented programming environment, but it should
be used with care to ensure that its benefits are maximized and its drawbacks are minimized.
