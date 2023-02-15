Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in the future. It means it lets you save and
restore the previous state of an object without revealing the details of its implementation.

The Memento design pattern usually has three actors:

1. Memento: A type that stores the type we want to save. Usually, we won't store the business type directly and we provide an extra layer of abstraction
   through this type.
2. Originator: A type that is in charge of creating mementos and storing the current active state. We said that the Memento type wraps states of the
   business type, and we use originator as the creator of mementos.
3. Care Taker: A type that stores the list of mementos that can have the logic to store them in a database or to not store more than a specified number
   of them.
