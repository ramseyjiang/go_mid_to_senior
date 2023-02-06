The Singleton pattern is one of the most used design patterns out there or one of the easiest to grasp. It will provide
a single instance of an object, and guarantee that there are no duplicates. At first, call to use the instance, if it's
created, and then reused between all the parts in the application that need to use the particular behavior.

Objectives:
1. Create a single, shared value, of some particular type.
2. Create a restricted object creation of some type to a single unit along the entire


Pros
1. You can be sure that a class has only a single instance.
2. You gain a global access point to that instance.
3. The singleton object is initialized only when it’s requested for the first time.


Cons
1. Violates the Single Responsibility Principle. The pattern solves two problems at the time.
2. The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each
   other.
3. It requires special solution in a multi-thread environment, if you use it directly in a multi-thread environment, it
   will be panic.
