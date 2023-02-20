Composite pattern is a structural design pattern that allows composing objects into a tree-like structure and working with it as if it was a singular
object. It lets you compose objects into tree structures and then work with these structures as if they were individual objects. The pattern uses a
single interface to represent both individual objects and groups of objects, which means that clients can interact with the objects without knowing
whether they are dealing with a single object or a group of objects.

Objective

1. Represents part-whole hierarchies of objects.
2. Allows clients to treat individual objects and groups of objects uniformly.
3. Simplifies the code by using a single interface for both individual objects and groups of objects.

How to implement

1. Make sure that the core model of your app can be represented as a tree structure. Try to break it down into simple elements and
   containers. Remember that containers must be able to contain both simple elements and other containers.
2. Declare the component interface with a list of methods that make sense for both simple and complex components.
3. Create a leaf class to represent simple elements. A program may have multiple different leaf classes.
4. Create a container class to represent complex elements. In this class, provide an array field for storing references to sub-elements. The
   array must be able to store both leaves and containers, so make sure it’s declared with the component interface type. While implementing
   the methods of the component interface, remember that a container is supposed to be delegating most of the work to sub-elements.
5. Finally, define the methods for adding and removal of child elements in the container.

Keep in mind these steps during you implement the component interface.

Pros

1. It allows you to treat individual objects and groups of objects uniformly, which makes the code easier to write and understand.
2. It simplifies the code by using a single interface for both individual objects and groups of objects.
3. It makes it easy to add new types of objects to the hierarchy.
4. It provides a flexible way to traverse the hierarchy of objects.

Cons

1. It can make the code more complex if not used properly.
2. It can be less efficient than other approaches if the hierarchy is very large.
3. It can be difficult to ensure that the composite structure is always in a valid state.

When using the Composite design pattern in Go, you must be very careful not to confuse it with inheritance. The inheritance problem is that every time
we want to add a new attribute, we would have to append it to the already defined classes. This is where composition comes in. The composition is always
easier to extend a new attribute.
