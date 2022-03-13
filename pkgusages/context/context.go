package contextpkg

import (
	"context"
	"fmt"
	"time"
)

// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// The chain of function calls between them must propagate the Context,
// optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.

/**
The interface of the Context is defined as:
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <- chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

Deadline: The first value is the deadline, at which point the Context will automatically trigger the Cancel action.
The second value is a boolean value, true means the deadline is set, false means the deadline is not set.
If the deadline is not set, you have to call the cancel function manually to cancel the Context.

Done: return a read-only channel (only after being canceled), type struct{}, when this channel is readable,
it means the parent context has initiated the cancel request, according to this signal,
the developer can do some cleanup actions, exit the goroutine

Err: returns the reason why the context was cancelled

Value: returns the value bound to the Context, it is a key-value pair,
so you need to pass a Key to get the corresponding value, this value is thread-safe.
*/

// Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
// The Context should be the first parameter, typically named ctx:
// func DoSomething(ctx context.Context, arg Arg) error {
// ... use ctx ...
// }

/*
	func Background() Context
	ctx, cancel:= context.Background()

	Background returns a non-nil, empty Context.
	It is never canceled, has no values, and has no deadline.
	It is typically used by the main function, initialization, and tests, and as the top-level Context
	of the tree structure, the root Context, which cannot be cancelled.
	It is used to derive other contexts and has an empty context as a return type.
*/

/*
	func TODO() Context
	ctx, cancel := context.TODO()

	TODO returns a non-nil, empty Context. Code should use context.
	TODO when it's unclear which Context to use or it is not yet available.
	because the surrounding function has not yet been extended to accept a Context parameter.

	TODO is to never pass nil context , instead, use of TODO is advised .
	TODO, when you don’t know what Context to use, you can use this.
*/

// both TODO() and Background() are essentially of type emptyCtx, both are non-cancelable, neither has a set deadline,
// and neither carries any value for the Context.

/*
	func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

	A CancelFunc tells an operation to abandon its work.
	A CancelFunc does not wait for the work to stop.
	A CancelFunc may be called by multiple goroutines simultaneously.
	After the first call, subsequent calls to a CancelFunc do nothing.

	When a Context is canceled, all Contexts derived from it are also canceled.
	The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent)
	and return a derived Context (the child) and a CancelFunc.
	Calling the CancelFunc cancels the child and its children,
	removes the parent's reference to the child, and stops any associated timers.

	For channel, although channel can also notify many nested goroutines to exit,
	channel is not thread-safe, while context is thread-safe.
*/

/**
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout and WithDeadline are basically the same in terms of usage and function,
they both indicate that the context will be automatically canceled after a certain time,
the only difference can be seen from the definition of the function,
the second parameter passed to WithDeadline is of type time.Duration type, which is a relative time,
meaning how long after the timeout is cancelled.
*/

/*
	func WithValue(parent Context, key, val interface{}) Context
	ctx := context.WithValue(context.Background(), key, “test”)

	WithValue returns a copy of parent in which the value associated with key is val.
	It creates a new context based on a parent context and adds a value to a given key.
	It allows you to store any type of data inside the context.
	Metadata is passed in as Key-Value, but note that the Key must be comparable and the Value must be thread-safe.
*/

//  Note that these methods mean that the context can be inherited once to achieve one more function,
//	for example, using the WithCancel function to pass in the root context, it creates a child context,
//	which has an additional function of cancel context, then use this context(context01) as the parent context,
//	and pass it as the first parameter to the WithDeadline function, the child context(context02) is obtained,
//	compared to the child context(context01), It has an additional function to cancel the context automatically
//	after the deadline.

/**
Context notes:

1. Do not store Contexts in struct types, but pass the Context explicitly to each function that needs it,
and the Context should be the first argument.

2. Do not pass a nil Context, even if the function allows it, or if you are not sure which Context to use,
pass context.

3. Do not pass variables that could be passed as function arguments to the Value of the Context.
*/

func Trigger() {
	withCancelUsage()
	withDeadlineUsage()
	withValueUsage()

	fmt.Println("---------------withCancelUsage2  start---------------")
	withCancelUsage2()
	fmt.Println("---------------withCancelUsage2  end-----------------")
	fmt.Println("---------------withTimeoutUsage  start---------------")
	withTimeoutUsage()
	fmt.Println("---------------withTimeoutUsage  end-----------------")
	fmt.Println("-------------withDeadlineUsage2  start---------------")
	withDeadlineUsage2()
	fmt.Println("-------------withDeadlineUsage2  end-----------------")
}

// This example demonstrates the use of a cancelable context to prevent a goroutine leak.
// By the end of the example function, the goroutine started by gen will return without leaking.
// WithCancel returns a copy of parent with a new Done channel.
// The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.

// context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithCancel(context.Background())
// This parent function is passed in as argument.
// This parent context can either be a background context or a context that was passed into the function.
func withCancelUsage() {
	// gen generates integers in a separate goroutine and sends them to the returned channel.
	// The callers of gen need to cancel the context once they are done consuming generated integers not to leak the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): // This branch is only reached when chan channel is closed, or when data is sent.
					fmt.Println("ctx.Done() is triggered, n=", n)
					return // returning not to leak the goroutine
				case dst <- n:
					fmt.Println("cancel usage is in progress: ", n)
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println("cancel usage: ", n)
		if n == 5 {
			break
		}
	}
}

// WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d.
// If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent.
// The returned context's Done channel is closed when the deadline expires, when the returned cancel function is called,
// or when the parent context's Done channel is closed, whichever happens first.
// When that context gets canceled because of the deadline running out, all the functions that got the context get notified to stop work and return.

// context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(2 * time.Second))

// Similar to WithDeadline , it cancels the contexts but this cancellation is based upon the time duration.
// This function returns a derived context that gets canceled if the cancel function is called or the timeout duration is exceeded.
// context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)
func withDeadlineUsage() {
	const shortDuration = 1 * time.Millisecond
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d) // ctx is Context type, cancel is CancelFunc type.

	// Even though ctx will be expired, it is good practice calling its cancellation function in any case.
	// Failure to do so may keep the context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("withDeadline usage:", ctx.Err())
	}
}

// This example demonstrates how a value can be passed to the context and also how to retrieve it if it exists.
func withValueUsage() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("Key is:", k, ", value is:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("color")

	// Define "PHP" is the value of k. It means key: color, value: PHP
	ctx := context.WithValue(context.Background(), k, "PHP")

	f(ctx, k)
	f(ctx, "test")
}

func monitor2(ctx context.Context, number int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("monitor2: %v, the received channel value is: %v, ending\n", number, v)
			return
		default:
			fmt.Printf("monitor2: %v in progress...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}
func monitor1(ctx context.Context, number int) {
	for {
		go monitor2(ctx, number)
		select {
		case v := <-ctx.Done():
			// this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
			fmt.Printf("monitor1: %v, the received channel value is: %v, ending\n", number, v)
			return
		default:
			fmt.Printf("monitor1: %v in progress...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func withCancelUsage2() {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	for i := 1; i <= 5; i++ {
		go monitor1(ctx, i)
	}
	time.Sleep(1 * time.Second)
	// close all gourtines
	cancel()
	// waiting 10 seconds, if the screen does not display <monitor: xxxx in progress>, all goroutines have been shut down
	time.Sleep(10 * time.Second)
}

func withDeadlineUsage2() {
	var ctx01 context.Context
	var ctx02 context.Context
	var cancelA context.CancelFunc
	var cancelB context.CancelFunc
	ctx01, cancelA = context.WithCancel(context.Background())

	// If it's WithTimeout, just change this line to "ctx02, cancel = context.WithTimeout(ctx01, 1 * time.Second)"
	ctx02, cancelB = context.WithDeadline(ctx01, time.Now().Add(1*time.Second))

	defer cancelA()
	defer cancelB()
	for i := 1; i <= 5; i++ {
		go monitor1(ctx02, i)
	}
	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("the cause of cancel is: ", ctx02.Err())
	}
}

func withTimeoutUsage() {
	var ctx01 context.Context
	var ctx02 context.Context
	var cancelA context.CancelFunc
	var cancelB context.CancelFunc

	ctx01, cancelA = context.WithCancel(context.Background())
	ctx02, cancelB = context.WithTimeout(ctx01, 1*time.Second)
	defer cancelA()
	defer cancelB()
	for i := 1; i <= 5; i++ {
		go monitor1(ctx02, i)
	}
	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("the cause of cancel is: ", ctx02.Err())
	}
}
