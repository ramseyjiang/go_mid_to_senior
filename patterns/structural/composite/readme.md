Composite pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.
In this pattern, you can treat all objects as the same via the common interface.

Pros

1. You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.
2. Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the
   object tree.

Cons It might be difficult to provide a common interface for classes whose functionality differs too much. In certain scenarios, you’d need
to overgeneralize the component interface, making it harder to comprehend.
