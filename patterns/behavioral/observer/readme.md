Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the
object they’re observing. That means the observer pattern allows some objects to notify other objects about changes in their state.

Objectives

1. Provides a way to decouple objects that are dependent on each other.
2. Changes the behavior of the Subject and the Observers independently, without affecting each other.

Pros

1. Decouples the Subject from the Observers, making the code more flexible and reusable.
2. Makes it easy to add new Observers to the Subject without affecting the existing code.
3. Promotes loose coupling, making the code more maintainable and testable.

Cons

1. Add performance overhead, especially if there are many Observers attached to a single Subject.
2. Add complexity to the code, making it harder to understand and maintain.
