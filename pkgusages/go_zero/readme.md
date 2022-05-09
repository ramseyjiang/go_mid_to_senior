go-zero is a web and rpc framework with lots of builtin engineering practices. 

In this simple sample, it manages multiple services in a process.


It’s born to ensure the stability of the busy services with resilience design and has been serving sites with tens of millions of users for years.

go-zero contains simple API description syntax and code generation tool called goctl. You can generate Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript from .api files with goctl.

Advantages of go-zero:

improve the stability of the services with tens of millions of daily active users
builtin chained timeout control, concurrency control, rate limit, adaptive circuit breaker, adaptive load shedding, even no configuration needed
builtin middlewares also can be integrated into your frameworks
simple API syntax, one command to generate a couple of different languages
auto validate the request parameters from clients
plenty of builtin microservice management and concurrent toolkits