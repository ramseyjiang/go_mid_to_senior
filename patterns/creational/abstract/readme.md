Abstract pattern is used to solve the problem of creating entire product families without specifying their concrete classes. It provides an
interface for creating all distinct products but leaves the actual product creation to concrete factory classes. Each factory type
corresponds to a certain product variety. Abstract pattern a new layer of grouping to achieve a bigger and more complex composite object,
which is used through its interfaces.

Objectives

1. Provide a new layer of encapsulation for Factory methods that return a common interface for all factories
2. Group common factories into a super Factory (also called a factory of factories)

Pros

1. You can be sure that the products you’re getting from a factory are compatible with each other.
2. You avoid tight coupling between concrete products and client code.
3. Single Responsibility Principle. You can extract the product creation code into one place, making the code easier to support.
4. Open/Closed Principle. You can introduce new variants of products without breaking existing client code.

Cons The code may become more complicated than it should be, since a lot of new interfaces and classes are introduced along with the pattern.

VehicleFactory is Abstract factory interface. It defines Build() method, but not implement it. Car is Concrete factory. Because they
implement Build(), and it defines NumDoors() also. Motorbike is Concrete factory. Because they implement Build(), and it defines
GetMotorbikeType() also. Vehicle is Abstract product, it defines NumWheels() and NumSeats(). FamilyCar and LuxuryCar are concrete products.
They implement NumWheels(), NumSeats() and NumDoors(). SportMotorbike and CruiseMotorbike are concrete products. They implement NumWheels(),
NumSeats() and GetMotorbikeType().
