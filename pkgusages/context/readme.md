What the context module does?
It’s an argument passed to your functions and Goroutines and allows you to stop them promptly should you not require
them anymore.

Imagine being the person taking orders in a restaurant.
When an order arrives, you delegate it to one of your many chefs.
What would you do if the customer suddenly decides to walk away?
Without a doubt, you will stop your chef from further processing the order to prevent any waste of ingredients!


What Is the Context Module?
Typical usage of the context module is when a client terminates the connection with a server.

What if the termination occurs while the server is in the middle of some heavy lifting work or database query?
The context module allows these processes to be stopped instantly as soon as they are not further in need.

The usage of the context module boils down to three primary parts
1. Listening to a cancellation event
2. Emitting a cancellation event
3. Passing request scope data


type Context interface {
// Channel listen to cancellation
Done() <-chan struct{}

// Return error if context is cancelled
Err() error

// Return the deadline set to the context
Deadline() (deadline time.Time, ok bool)

// Return the value for a given key
Value(key interface{}) interface{}

}

// Return empty root context
func Background() Context
func TODO() Context

// Return cancellable context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

// Return context with key-value pairs
func WithValue(parent Context, key, val interface{}) Context


type Context interface {
Done() <- chan struct{}
Err() error

Deadline() (deadline time.Time, ok bool)
Value(key interface{}) interface{}

}

The Context type is nothing but an interface that implements four simple functions.

Done: returns a read-only channel (only after being canceled), type struct{}, when this channel is readable,
it means the parent context has initiated the cancel request, according to this signal,
the developer can do some cleanup actions, exit the goroutine

Err: returns a non-nil error in the event of cancellation otherwise, it returns a nil value.

Deadline: The first value is the deadline, at which point the Context will automatically trigger the Cancel action.
The second value is a boolean value, true means the deadline is set, false means the deadline is not set.
If the deadline is not set, you have to call the cancel function manually to cancel the Context.

Value: returns the value bound to the Context, it is a key-value pair,
so you need to pass a Key to get the corresponding value, this value is thread-safe.

Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
The Context should be the first parameter, typically named ctx:
func DoSomething(ctx context.Context, arg Arg) error {
... use ctx ...

}

The context module provides three functions that return a CancelFunc .
Calling the cancelFunc emits an empty struct to the ctx.Done() channel and notifies downstream functions that are listening to it.

// Return cancellable context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

The WithCancel function takes in a parent context and returns a cancellable context and a cancel function.
The WithTimeout allows you to specify a timeout duration and automatically cancels the context if the duration exceeds.
The WithDeadline function accepts a specific timeout time instead of a duration. Other than that, it works exactly similar as WithTimeout


func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

When a Context is canceled, all Contexts derived from it are also canceled.
WithCancel, WithDeadline, WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc.
Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers.


func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout and WithDeadline are basically the same in terms of usage and function,
they both indicate that the context will be automatically canceled after a certain time,
the only difference can be seen from the definition of the function,
the second parameter passed to WithDeadline is of type time.Duration type, which is a relative time,
meaning how long after the timeout is cancelled.


// Return context with key-value pairs
func WithValue(parent Context, key, val interface{}) Context
ctx := context.WithValue(context.Background(), key, “test”)

WithValue returns a copy of parent in which the value associated with key is val.
It creates a new context based on a parent context and adds a value to a given key.
It allows you to store any type of data inside the context.
Metadata is passed in as Key-Value, but note that the Key must be comparable and the Value must be thread-safe.


As you call the WithX functions, they accept a parent context and return a new copy of the parent with a new Done channel.

rootCtx := context.Background()
child1Ctx, cancelFunc1 := context.WithCancel(rootCtx)
child2Ctx, cancelFunc2 := context.WithCancel(rootCtx)

child3Ctx, cancelFunc3 := context.WithCancel(child1Ctx)

Since the functions require a parent context as an argument,
the context module offers two simple functions to create a root context.


// Return empty root context
func Background() Context
func TODO() Context


These functions output an empty context that does nothing AT ALL. It cannot be cancelled nor carry a value.
Their primary purpose is to serve as a root context that will later be passed to any of the WithX functions to create a cancellable context.

Note that these methods mean that the context can be inherited once to achieve one more function,
for example, using the WithCancel function to pass in the root context, it creates a child context,
which has an additional function of cancel context, then use this context(context01) as the parent context,
and pass it as the first parameter to the WithDeadline function, the child context(context02) is obtained,
compared to the child context(context01), It has an additional function to cancel the context automatically
after the deadline.

Context notes:

1. Do not store Contexts in struct types, but pass the Context explicitly to each function that needs it,
   and the Context should be the first argument.

2. Do not pass a nil Context, even if the function allows it, or if you are not sure which Context to use, pass context.

3. Do not pass variables that could be passed as function arguments to the Value of the Context.

4. Remember using WithCancel in your code, except it has WithTimeout and WithDeadline.

When you spawn a new cancellable context via the WithCancel function, the module will:
a, Spawn a new Goroutine in the background to propagate the cancellation event to all children if the cancel function is invoked
b, Keep track of all the children contexts in the struct of the parent context

If a function returns without cancelling the context, the Goroutine and the child contexts will remain in the memory
indefinitely causing a memory leak.

This also applies to WithTimeout and WithDeadline except, these functions automatically cancel the context when the deadline is exceeded.