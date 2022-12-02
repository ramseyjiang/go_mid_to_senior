In Golang that implements 3 basic rate limiting algorithms:

1. Throttle Rate Limiter Limits based on a single request per specified time interaval. A throttler is to handle burst
   programs by only allowing a certain number of requests to be processed per time duration. For example, a public API
   may try to regulate the server load by only allowing 1 request per second, per client.

2. Max Concurrency Rate Limiter Limits the number of active concurrent requests at any given time. If a rate limit is
   requested and the active count is less than the limit, the rate limit is immediately granted, otherwise the request
   must wait until a previously granted rate limit is finished. This means that we need to add the ability to release a
   rate limit token.

3. Fixed Window Rate Limiter Specifies a fixed number of requests that can be processed in a given time window. Once
   that limit is reached no more requests can be processed until the next window. Each request that occurs within the
   window time increments the counter until the limit is reached, then each subsequent request has to wait until the
   next available window.

Rate limit should follow these steps below:

1. The rate limit should expose an Acquire() method that, when called, will block until a rate limit token is available.
2. Interally, we’ll have two separate channels for synchronization: 1) an incoming request channel, and 2) an outgoing
   token channel.
3. Each Rate Limiter will define an await function that listens for incoming requests, determines when a rate limit
   token is available, and sends the token to the out channel.