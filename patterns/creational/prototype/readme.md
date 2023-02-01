The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface
makes the client code independent of the concrete classes of objects that it clones. The aim of the Prototype pattern is to have an object or
a set of objects that is already created at compilation time, but you can clone as many times as you want at runtime.An object that supports
cloning is called a prototype. Here’s how it works: you create a set of objects, configured in various ways. When you need an object like the
one you’ve configured, you just clone a prototype instead of constructing a new object from scratch.

Objective

1. Avoid repetitive object creation.
2. Maintain a set of objects that will be cloned to create new instances
3. Provide a default value of some type to start working on top of it
4. Free CPU of complex object initialization to take more memory resources

Pros

1. You can clone objects without coupling to their concrete classes.
2. You can get rid of repeated initialization code in favor of cloning pre-built prototypes.
3. You can produce complex objects more conveniently.
4. You get an alternative to inheritance when dealing with configuration presets for complex objects.

Cons Cloning complex objects that have circular references might be very tricky.
