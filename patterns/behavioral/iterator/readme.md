The Iterator pattern is a behavioral design pattern that provides a way to traverse a collection of objects without exposing the underlying
implementation details. It allows you to access the elements of a collection sequentially, one at a time, without having to know the internal structure
of the collection.

Objectives

1. Provides a standardized way of iterating over a collection of objects
2. helps to decouple the iteration logic from the collection, which improves code readability, maintainability, and reusability.

Pros

1. Separation of concerns: The Iterator separates the logic of accessing the elements of a collection from the collection itself, which leads to
   cleaner, more modular code.
2. High reusable: Once an Iterator is implemented, it can be used to iterate over any collection that implements the Iterator interface.
3. Encapsulation: The Iterator provides a way to encapsulate the internal representation of a collection, which makes it easier to change the collection
   without affecting the code that uses the Iterator.

Cons

1. Overhead: it may introduce some overhead, as it requires an additional object to be created to maintain the iteration state.
2. Inefficient for small collections: For very small collections, the overhead of using an Iterator may outweigh the benefits.

How to implement

The Iterator pattern can be implemented using an interface that defines the methods for iterating over a collection. The interface typically includes a
Next() method that returns the next element of the collection, and a HasNext() method that checks whether there are more elements in the collection.
