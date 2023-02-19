Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the
object they’re observing. That means the observer pattern allows some objects to notify other objects about changes in their state. It is also known as
publish/subscriber or publish/listener design pattern.

Objectives

1. Provides a way to decouple objects that are dependent on each other.
2. Changes the behavior of the Subject and the Observers independently, without affecting each other.

Pros

1. Provide an event-driven architecture where one event can trigger one or more actions
2. Uncouple the actions that are performed from the event that triggers them
3. Provide more than one event that triggers the same action

Cons

1. Add performance overhead, especially if there are many Observers attached to a single Subject.
2. Add complexity to the code, making it harder to understand and maintain.

The Observer pattern is especially useful to achieve many actions that are triggered on one event. It is also especially useful when you don't know how
many actions are performed after an event in advance or there is a possibility that the number of actions is going to grow in the near future.
Event-driven architectures includes the State pattern and the Observer pattern.
