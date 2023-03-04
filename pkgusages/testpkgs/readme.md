1. t.Fail(): This marks the test as failed, but allows the execution to continue.

2. t.FailNow(): marks the test as failed and stops its execution immediately using runtime.Goexit().


3. t.Errorf(): This combines logging an error message t.Logf() with failing the test t.Fail().

If you have multiple test cases in a test function that do not depend on each other, and you do not want to halt execution but instead mark the test as
failed if any of them fail, you can use t.Errorf()

4. t.Fatalf(): combines logging an error message using t.Logf() with immediately failing the test using t.FailNow().

If you have a series of test cases that rely on each other, and you want to stop the execution immediately if any one of them fails, then you may want
to consider a different approach using t.Fatalf()

5. t.Parallel() to turn on the parallel mode

To run subtests in parallel, use t.Parallel() to turn on the parallel mode. This can be useful when test cases are independent of each other, since it
can make our tests run faster
