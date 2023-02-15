Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in the future. It means it lets you save and
restore the previous state of an object without revealing the details of its implementation.

The Memento pattern is a behavioral design pattern that allows an object to save its internal state so that it can be restored later, without violating
encapsulation. In other words, it provides a way to capture the current state of an object and restore it to that state later on.

The Memento design pattern usually has three actors:

1. Memento: A type that stores the type we want to save. Usually, we won't store the business type directly and we provide an extra layer of abstraction
   through this type.
2. Originator: A type that is in charge of creating mementos and storing the current active state. We said that the Memento type wraps states of the
   business type, and we use originator as the creator of mementos.
3. Care Taker: A type that stores the list of mementos that can have the logic to store them in a database or to not store more than a specified number
   of them.

Objectives

1. To provide a way to store the internal state of an object, without exposing its implementation details.
2. To allow an object to restore its previous state.
3. To make the restoration of the object's state undoable.

Pros:

1. It provides a way to store an object's state without exposing its implementation details, thus adhering to the principle of encapsulation.
2. It makes it easy to restore an object's state to a previous state.
3. It allows you to implement undo/redo functionality easily.
4. It separates the responsibility of storing and restoring an object's state, making the code easier to maintain.

Cons:

1. It can lead to a large number of objects being created if the state changes frequently.
2. It can impact the performance of the system if the state being stored is large.
