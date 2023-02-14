The Chain of Responsibility pattern is a design pattern that allows multiple objects to handle a request, with each object having the opportunity to
either handle the request or pass it along to the next object in the chain. The objects in the chain are not aware of each other, but instead
communicate through a common interface, allowing new objects to be added or removed from the chain easily.

Objectives

1. To decouple the sender of a request from its receiver.
2. To allow for the dynamic modification of the chain of objects that handle requests.
3. To allow multiple objects to handle a request, with each object having the opportunity to either handle the request or pass it along to the next
   object in the chain.

Pros

1. Loose coupling: The objects in the chain are decoupled from each other, so changes in one object do not affect the others.
2. Increased flexibility: New objects can be added to the chain or existing objects can be removed, making it easy to change the behavior of the chain.
3. Improved maintainability: The logic of handling requests can be changed easily by changing the order of objects in the chain or adding/removing
   objects from the chain.

Cons

1. Performance overhead: If there are many objects in the chain, passing requests through them can become slow and resource-intensive.
2. Increased complexity: Maintaining the chain can become complex as the number of objects in the chain increases.
3. Debugging difficulties: Lead to result in a long chain of objects, making it difficult to understand how a request is being handled and making it
   harder to debug problems.

How to implement

1. Define an interface that defines the methods for handling requests.
2. Create concrete objects that implement the interface.
3. Each object would have a reference to the next object in the chain, and would use that reference to pass along requests that it could not handle.

The Chain of Responsibility pattern opens the door of middleware of any type and plugin like libraries to improve the functionality of some part. Many
open source projects uses a Chain of Responsibility to handler HTTP requests and responses to extract information to the end user (such as cookies info)
or check authentication details (I'll let you pass to the next link only if I have you on my database).

In summary, the Chain of Responsibility pattern is a powerful tool for processing requests in an object-oriented programming environment, but it should
be used with care to ensure that its benefits are maximized and its drawbacks are minimized.
