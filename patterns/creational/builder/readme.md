The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic
they require. It will reuse an algorithm to create many implementations of an interface. It means the builder pattern
will allow you to produce different types and representations of an object using the same construction code. The builder
pattern's instance will provide the opening and closing braces {} and leaving the instance with zero values, or as
complex as an object that needs to make some API calls, check states, and create objects for its fields. The
BuildProcess interface specifies what it must comply to be part of the possible builders.

Objectives
1. Abstract complex creations so that object creation is separated from the object user
2. Create an object step by step by filling its fields and creating the embedded objects
3. Reuse the object creation algorithm between many objects


Pros
1. You can construct objects step-by-step, defer construction steps or run steps recursively.
2. You can reuse the same construction code when building various representations of products.
3. Single Responsibility Principle. You can isolate complex construction code from the business logic of the product.

Cons
1. The overall complexity of the code increases since the pattern requires creating multiple new classes.
