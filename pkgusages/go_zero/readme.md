go-zero is a web and rpc framework with lots of builtin engineering practices. 

In this simple sample, it manages multiple services in a process.

**Using "curl -i http://localhost:8080/morning" and "curl -i http://localhost:8080/evening" in command to test outputs.**

It’s born to ensure the stability of the busy services with resilience design and has been serving sites with tens of millions of users for years.

go-zero contains simple API description syntax and code generation tool called goctl. You can generate Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript from .api files with goctl.

Advantages of go-zero:

improve the stability of the services with tens of millions of daily active users
builtin chained timeout control, concurrency control, rate limit, adaptive circuit breaker, adaptive load shedding, even no configuration needed
builtin middlewares also can be integrated into your frameworks
simple API syntax, one command to generate a couple of different languages
auto validate the request parameters from clients
plenty of builtin microservice management and concurrent toolkits


Using "brew install cloc" to install the cloc tool.
cloc is used to count the blank lines, comment lines, actual code lines written in many programming languages.
Github: https://github.com/AlDanial/cloc/

% pwd
/Users/daweijiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.3.3

daweijiang@DAWEIs-MBP go-zero@v1.3.3 % cloc core/service/servicegroup.go
1 text file.
1 unique file.                              
0 files ignored.

github.com/AlDanial/cloc v 1.92  T=0.01 s (195.3 files/s, 23042.4 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               1             22             14             82
-------------------------------------------------------------------------------