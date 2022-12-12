5 points make write better codes for tests.

1. Apply clean code: write short functions, handle a single task per function, etc.
2. Write extendable code through the use of abstractions, interfaces and mocks.
3. Understand the business logic better by testing regular/edge cases and high coverage of these.
4. Avoid legacy, long-untouched and/or unmaintainable code — tests will ease the process of maintaining changes to code,
   so it does not rot.
5. Measure the performance of your code through benchmarks, load tests, etc.

5 Mocking Techniques:

1. Higher-Order Functions
2. Monkey Patching
3. Interface Substitution
4. Embedding Interfaces
5. Mocking out Downstream HTTP Calls with net/http/httptest && GRPC response

These 5 mocking techs are almost matching 5 points write better codes for tests.

assert.Nil() and require.Nil(), both are used to check err is nil or not. require.Nil() always is used for checking
requests have all params or not.