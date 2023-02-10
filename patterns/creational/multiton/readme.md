The Multiton pattern is a creational design pattern that is similar to the Singleton pattern, but with a twist. The Singleton pattern ensures that a
class has only one instance, while the Multiton pattern ensures that a class has a limited number of instances, each with its own unique identifier. The
Multiton pattern is useful in situations where multiple instances of a class are required, but each instance must have a unique identity.

Objectives:

To ensure that a limited number of instances of a class are created and each instance has a unique identifier. To provide a common point of access to
all instances of the class, allowing them to be easily managed and accessed.

Pros:

Enforces the limit on the number of instances that can be created, helping to avoid resource exhaustion and memory leaks. Provides a common point of
access to all instances of the class, making it easy to manage and access the instances. Supports the creation of different configurations for each
instance, making it possible to use the Multiton pattern in a wide range of applications.

Cons:

Can result in increased complexity, as the code to manage and access the instances must be carefully designed and implemented. May result in decreased
performance, as the code to manage the instances may take longer to run than if there were only one instance of the class.
