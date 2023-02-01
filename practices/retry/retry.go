package retry

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

/**
User story:
Let’s say the company you work for develops a to-do list manager (yes, not very imaginative).
This app has a feature where users receive an email with a PDF report summarising their completed tasks.
A third-party service handles the generation of these PDFs.
Sadly the service is often plagued by transient faults.
A transient fault, also known as a transient error, has an underlying cause that resolves itself.
The third-party service is down now and again but always recovers quickly.
You are tasked with making this service more reliable since the users seem to be quite fond of it.
Luckily the retry pattern offers a straightforward but effective solution to your woes.
*/

var count int

func Entry() {
	r := attempt(getPdfURL, 5, 2*time.Second)
	res, err := r(context.Background())
	fmt.Println(res, err)
}

// getPdfURL is a function which the signature of this function needs to match the Effector type.
func getPdfURL(ctx context.Context) (string, error) {
	count++
	if count <= 3 {
		return "", errors.New("boom")
	} else {
		return "https://linktopdf.com", nil
	}
}

// Effector type defines a function signature, and this function interacts with the third party service.
type Effector func(context.Context) (string, error)

// attempt accepts an Effector and returns an anonymous function with the same signature as the Effector.
// It essentially wraps the received function with the retry logic.
// The function accepts three parameters: an Effector, an integer describing how many times the function retries the passed Effector
// and the delay between the retries.
// The Retry pattern, some form of backoff algorithm is implemented that increases the delay between each retry.
func attempt(effector Effector, retries int, delay time.Duration) Effector {
	// Declare a function is to use a function literal.
	// A function literal is written like a function declaration without a name following the func keyword.
	// The value of this expression is called an anonymous function.
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				// Return when there is no error or the maximum amount
				// of retries is reached.
				return response, err
			}

			log.Printf("Function call failed, retrying in %v", delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
