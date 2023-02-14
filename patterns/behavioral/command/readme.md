Command is a behavioral design pattern that allows encapsulating a request or operation as an object. It converts requests or simple operations into a
stand-alone object that can be stored, passed, and executed independently of the original requester. The pattern separates the requester (client) from
the object that performs the action (receiver). This transformation lets you pass requests as a method arguments, delay or queue a request’s execution,
and support undoable operations.

The Command Pattern aims to achieve the following objectives:

1. Encapsulate a request as an object
2. Decouple the client from the receiver
3. Allow requests to be queued, logged, or undone
4. Enable implementing a transactional system

Pros:

1. It provides loose coupling between the client and the receiver.
2. It enables logging, undo, and redo operations.
3. It allows queuing requests and processing them at a later time.
4. It facilitates implementing a transactional system.

Cons:

1. It can create a large number of command classes, which may lead to increased complexity.
2. It may impact performance, particularly when executing a large number of commands.

How to implement in Golang:

1. Define the Command interface that declares the common method for all concrete commands.
2. Create concrete command structs that implement the Command interface.
3. Define the Receiver interface that declares the methods that the command object calls.
4. Create one or more Receiver implementations.
5. Create an Invoker struct that has a method to execute the command.
6. Instantiate the Invoker, Command, and Receiver objects, and set the receiver as the target of the command.
7. Invoke the command using the Invoker, which will call the command's execute method, which will delegate the call to the Receiver.
