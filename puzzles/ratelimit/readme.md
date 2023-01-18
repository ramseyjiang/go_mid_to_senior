**WHat is rate limiter is used to do?**
In network, rate limiting is used to limit the rate of the requests. Rate limiter is widely used in production to
protect underlying services and resources by limiting the number of client requests allowed to be sent over a specified
period.

Pros:
Prevent resource starvation caused by Denial of Service (DoS) attack. Prevent servers from being overloaded. Controlling
flow.

In Golang that implements 3 basic rate limiting algorithms:

1. Leaky bucket Algorithm The bucket is like a queue or buffer, requests are processed at a fixed rate. Requests will be
   added to the bucket as long as the bucket is not full, any extra request spills over the bucket edge are discarded.

   Leaky bucket Algorithm Limits based on a single request per specified time interval. A throttler is to handle burst
   programs by only allowing a certain number of requests to be processed per time duration. For example, a public API
   may try to regulate the server load by only allowing 10 requests per second, per client.

2. Token bucket Algorithm The bucket has a capacity of tokens and the tokens will be refilled at some rate. Each request
   will attempt to withdraw a token from the bucket, if there are no tokens in the bucket, the service has reached its
   limit, otherwise, the request goes through.

   Max Concurrency Rate Limiter Limits the number of active concurrent requests at any given time. If a rate limit is
   requested and the active count is less than the limit, the rate limit is immediately granted, otherwise the request
   must wait until a previously granted rate limit is finished. This means that we need to add the ability to release a
   rate limit token.

3. Fixed window Algorithm Windows are split upfront and each window has a counter. Each request increases the counter by
   one. Once the counter reaches the threshold, new requests are dropped until a new time window starts. This algorithm
   is easy to implement, but they are subject to spike at the edges of the window.

   Fixed Window Rate Limiter Specifies a fixed number of requests that can be processed in a given time window. Once
   that limit is reached no more requests can be processed until the next window. Each request that occurs within the
   window time increments the counter until the limit is reached, then each subsequent request has to wait until the
   next available window.

Rate limit should follow these steps below:

1. The rate limit should expose Acquire() method that, when called, will block until a rate limit token is available.
2. Internally, we’ll have two separate channels for synchronization:
   1) an incoming request channel,
   2) an outgoing token channel.
3. Each Rate Limiter will define an await function that listens for incoming requests, determines when a rate limit
   token is available, and sends the token to the out channel.

Folder read guide

1. Define RateLimiter interface and config struct in the ratelimit.go file.
2. Define a Manager struct that implements a Rate Limiter Interface. In channel and out channel fields are in the
   manager. In the manager, it also includes a field named makeToken, which is a factory function for creating tokens
   that will allow different rate limiter implementations to define their own custom logic for token creation.
