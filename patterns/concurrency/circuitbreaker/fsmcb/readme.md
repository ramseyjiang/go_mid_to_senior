The circuit breaker. It is a mechanism that allows you to protect your service from performing too many requests in a
short period.

The circuit breaker has three distinct states:
Closed — all the requests are allowed to pass through to the upstream service. Open — All the requests are not allowed
to pass through to the upstream service. Half-Open — To determine if the upstream service recovered, the circuit breaker
will allow only a small number of requests to pass through.

How the circuit breaker state changes? Close to Open — When failed requests exceed the threshold in the close state, it
will change to the open state. Open to Half-Open — When a certain timeout period passes in the open state, it will
change to the half-open state. Half-Open to Open — When a request to upstream service still fails in the half-open
state, it will change to the open state again. Half-Open to Closed — When a certain number of predefined requests are
successful in a half-open state, it will change to the closed state.

The states change picture is following the link: https://miro.medium.com/max/1400/0*NL1JUDCDs_HA947e.png.

Server Have an HTTP server running on port 8080 as the upstream service. To simulate the upstream service is down, we
return a 500 error code to the client in the first 5 seconds on startup.

Client Define a simple function to call the upstream service on the client side.

main.go A circuit breaker with its configuration:
Name is the name of the circuit breaker MaxRequests is the maximum number of requests allowed to pass through when the
circuit breaker is half-open. Interval is the cyclic period of the closed state for the circuit breaker to clear the
internal Counts. Timeout is the period of the open state, after which the state of the circuit breaker becomes
half-open. ReadyToTrip is called whenever a request fails in the closed state. The circuit breaker will come into the
open state if ReadyToTrip returns true. OnStateChange is called whenever the state of the CircuitBreaker changes.

In the first second, the circuit breaker found that the upstream service consecutively failed more than three times; it
switched to the open state from the closed state. After the timeout period passes, the circuit breaker will switch to
the half-open state. The circuit breaker switches to an open state when the following request return error from the
upstream service. When upstream recovers, the circuit breaker observes this change by making three continuous successful
requests. Then it switches to the closed state.

An interesting question. If I put all codes in a file, then it can output the circuit breaker state. But if I separate
codes to three files, such as client.go, server.go and main.go. After that, it cannot output the circuit breaker state.