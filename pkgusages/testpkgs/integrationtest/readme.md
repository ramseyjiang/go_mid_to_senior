Integration Testing:
Applications consist of multiple modules, the interaction between those modules will give users outputs and results they
want.

an Integration test test means when a programmer wants to test this interaction and check if the data flow is working as
expected and giving the expected output, this is called.

Integration test will cover the interaction of those dependencies so there will be no mocking and less fake data, and as
I realized and practiced only the third party can be mocked.

Unit Test:

Unit tests are more about testing the behavior of a single unit (function, method, or certain piece of code). In this
type of testing, programmers will use fake data and mocking objects to define and expect the behavior of the unit
dependencies. The result of it will be asserting the expected result of this unit.

Integration Test vs Unit Test:
Furthermore, since the Integration test is made to test the whole workflow we have to assert database insertion and
other transactions. Also, we have to assert other storage behavior (S3 and Redis) if they are used. All of that will
make the Integration test more complicated and might fail easily if any module changes, so it has to be maintained more
often than the Unit test.

The Integration test is like the container for the Unit test where all the small units will be tested indirectly inside
it.