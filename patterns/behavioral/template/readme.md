The Template pattern is a behavioral design pattern that defines the skeleton of an algorithm in a base class and allows subclasses to override some
steps of the algorithm without changing its structure. In other words, it provides a way to define the steps of an algorithm, while allowing subclasses
to implement specific behaviors that can vary.

Pros

1. Code reuse: Promotes code reuse by providing a common structure for related algorithms.
2. Flexibility: Allows subclasses to implement specific behaviors that can vary, making it a flexible design pattern.
3. Maintenance: Makes it easier to maintain code by centralizing the algorithm's steps in a single place.

Cons:

1. Inflexibility: Be inflexible if the steps of the algorithm cannot be easily generalized.
2. Complexity: Be more complex than other design patterns, which can make it harder to understand and implement.

How to implement

1. To define the steps of an algorithm in a base class.
2. To allow subclasses to implement specific behaviors that can vary.
3. To avoid code duplication by providing a common code structure.
