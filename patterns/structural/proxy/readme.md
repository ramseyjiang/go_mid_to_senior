Proxy is a structural design pattern that provides an object that acts as a substitute or placeholder for a real service object used by a client. A
proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object. That means a proxy
controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.

The Proxy pattern usually wraps an object to hide some of its characteristics.

1. Lazy initialization (virtual proxy). This is when you have a heavyweight service object that wastes system resources by being always up, even though
   you only need it from time to time. Instead of creating the object when the app launches, you can delay the object’s initialization to a time when
   it’s really needed.
2. Access control (protection proxy). This is when you want only specific clients to be able to use the service object; for instance, when your objects
   are crucial parts of an operating system and clients are various launched applications (including malicious ones). The proxy can pass the request to
   the service object only if the client’s credentials match some criteria.
3. Remote proxy. Local execution of a remote service. This is when the service object is located on a remote server. In this case, the proxy passes the
   client request over the network, handling all the nasty details of working with the network.
4. Logging requests (logging proxy). This is when you want to keep a history of requests to the service object. The proxy can log each request before
   passing it to the service.
5. Caching proxy. It is used to cache request results. This is when you need to cache results of client requests and manage the life cycle of this
   cache, especially if results are quite large. The proxy can implement caching for recurring requests that always yield the same results. The proxy
   may use the parameters of requests as the cache keys.

Objectives

1. Hide an object behind the proxy so the features can be hidden, restricted, and so on.
2. Provide a new abstraction layer that is easy to work with, and can be changed easily.

How to Implement

1. If there’s no pre-existing service interface, create one to make proxy and service objects interchangeable. Extracting the interface from the service
   class isn’t always possible, because you’d need to change all the service’s clients to use that interface. Plan B is to make the proxy a subclass of
   the service class, and this way it’ll inherit the interface of the service.
2. Create the proxy class. It should have a field for storing a reference to the service. Usually, proxies create and manage the whole life cycle of
   their services. On rare occasions, a service is passed to the proxy via a constructor by the client.
3. Implement the proxy methods according to their purposes.
4. Consider introducing a creation method that decides whether the client gets a proxy or a real service. This can be a simple static method in the
   proxy class or a full-blown factory method.
5. Consider implementing lazy initialization for the service object.

Pros

1. You can control the service object without clients knowing about it.
2. You can manage the lifecycle of the service object when clients don’t care about it.
3. The proxy works even if the service object isn’t ready or is not available.
4. Open/Closed Principle. You can introduce new proxies without changing the service or clients.

Cons

1. The code may become more complicated since you need to introduce a lot of new classes.
2. The response from the service might get delayed.

The Proxy must implement the same interfaces as the features of the type it tries to wrap
