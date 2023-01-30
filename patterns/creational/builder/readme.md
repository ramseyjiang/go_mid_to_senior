The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic
they require. It will reuse an algorithm to create many implementations of an interface. It means the builder pattern
will allow you to produce different types and representations of an object using the same construction code. The builder
pattern's instance will provide the opening and closing braces {} and leaving the instance with zero values, or as
complex as an object that needs to make some API calls, check states, and create objects for its fields.

Requirements and acceptance criteria:

1. Must have a manufacturing type that constructs everything that a vehicle needs.
2. When using a car builder, the VehicleProduct with four wheels, five seats, and a structure defined as Car must be
   returned.
3. When using a bike builder, the VehicleProduct with two wheels, 1 seats, and a structure defined as Bike must be
   returned
4. A VehicleProduct built by any BuildProcess builder must be open to modifications.

Pros

1. You can construct objects step-by-step, defer construction steps or run steps recursively.
2. You can reuse the same construction code when building various representations of products.
3. Single Responsibility Principle. You can isolate complex construction code from the business logic of the product.

Cons The overall complexity of the code increases since the pattern requires creating multiple new classes.
