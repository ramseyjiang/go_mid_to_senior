package contextpkg

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const InternalSeconds = 5

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

	_ = http.ListenAndServe(":8000", http.HandlerFunc(withDoneUsage))
}

/**
handlerDone is used time.After() to simulate a function that takes five seconds to process a request.

If the context is cancelled within two seconds, the ctx.Done() channel receives an empty struct.
The second case will be executed and the function exits.

You can fire up this code on your local.
Once up, visit localhost:8000 on your browser, then close it within two seconds.
Observe your terminal and see what happens.
*/
func withDoneUsage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(InternalSeconds * time.Second):
		_, err := w.Write([]byte("request processed"))
		if err != nil {
			return
		}
	case <-ctx.Done():
		fmt.Println("request cancelled")
		return
	}
}

// This example demonstrates the use of a cancelable context to prevent a goroutine leak.
// By the end of the example function, the goroutine started by gen will return without leaking.
// WithCancel returns a copy of parent with a new Done channel.
// The returned context's Done channel is closed when the returned cancel function is called or
// when the parent context's Done channel is closed, whichever happens first.

// context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// ctx, cancel := context.WithCancel(context.Background())
// This parent function is passed in as argument.
// This parent context can either be a background context or a context that was passed into the function.
func withCancelUsage() {
	// gen generates integers in a separate goroutine and sends them to the returned channel.
	// The callers of gen need to cancel the context once they are done consuming generated integers not to leak the internal goroutine.
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
