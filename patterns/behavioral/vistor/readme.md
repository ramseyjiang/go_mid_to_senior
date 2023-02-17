Visitor is a behavioral design pattern that allows you to add new algorithms from the objects structure on which they operate. It involves defining a
separate object (the visitor) that can visit each object in a complex structure and perform some operation on it. This way, you can modify the behavior
of a group of objects without changing their individual classes.

Objectives

1. Separate the algorithm from the object structure it operates on.
2. Add new operations to a group of classes without modifying them individually.
3. Maintain the open-closed principle, which states that software entities should be open for extension but closed for modification.

Pros

1. Allows you to add new operations without modifying existing code.
2. Encourages the use of inheritance and polymorphism, which can make code more flexible and reusable.
3. Encapsulates the visitor algorithm in a separate object, making it easier to test and maintain.
4. Can be used with complex object structures that are difficult to modify.

Cons

1. Make code more complex by introducing additional objects and interfaces.
2. Be difficult to implement correctly, especially when dealing with complex object structures.
3. Lead the result in a proliferation of visitor classes if you need to perform many operations.
