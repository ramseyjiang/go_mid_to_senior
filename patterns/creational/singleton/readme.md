Objectives When we consider using the Singleton pattern we should follow rules following:

1. Create a single, shared value, of some particular type.
2. Create a restrict object creation of some type to a single unit along the entire

Requirements and acceptance test criteria:

1. When no counter has been created before, a new one is created with the value 0.
2. If a counter has already been created, return this instance that holds the actual count.
3. If we call the method AddOne, the count must be incremented by 1.

Pros

1. You can be sure that a class has only a single instance.
2. You gain a global access point to that instance.
3. The singleton object is initialized only when it’s requested for the first time.

Cons

1. Violates the Single Responsibility Principle. The pattern solves two problems at the time.
2. The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each
   other.
3. The pattern requires special treatment in a multi-thread environment so that multiple threads won’t create a
   singleton object several times.
