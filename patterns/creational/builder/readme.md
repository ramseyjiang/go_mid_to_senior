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
