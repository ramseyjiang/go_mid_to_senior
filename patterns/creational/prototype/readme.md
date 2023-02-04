The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface makes the
client code independent of the concrete classes of objects that it clones. The aim of the Prototype pattern is to have an object or a set of objects
that is already created at compilation time, but you can clone as many times as you want at runtime. An object that supports cloning is called a
prototype. Here’s how it works: you create a set of objects, configured in various ways. When you need an object like the one you’ve configured, you
just clone a prototype instead of constructing a new object from scratch.

Objective

1. Avoid repetitive object creation.
2. Maintain a set of objects that will be cloned to create new instances
3. Provide a default value of some type to start working on top of it
4. Free CPU of complex object initialization to take more memory resources

How to Implement

1. Create the prototype interface and declare the clone method in it. Or just add the method to all classes of an existing class hierarchy, if you
   have one.
2. A prototype class must define the alternative constructor that accepts an object of that class as an argument. The constructor must copy the values
   of all fields defined in the class from the passed object into the newly created instance. If you’re changing a subclass, you must call the parent
   constructor to let the superclass handle the cloning of its private fields.
3. The cloning method usually consists of just one line: running a new operator with the prototypical version of the constructor. Note, that every
   class must explicitly override the cloning method and use its own class name along with the new operator. Otherwise, the cloning method may produce
   an object of a parent class.
4. Optionally, create a centralized prototype registry to store a catalog of frequently used prototypes.

Pros

1. You can clone objects without coupling to their concrete classes.
2. You can get rid of repeated initialization code in favor of cloning pre-built prototypes.
3. You can produce complex objects more conveniently.
4. You get an alternative to inheritance when dealing with configuration presets for complex objects.

Cons Cloning complex objects that have circular references might be very tricky.
