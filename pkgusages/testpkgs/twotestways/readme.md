Each doughnut-box has its capacity, if doughnuts number is larger than the box capacity, it will return error. Doughnuts
have five different types. If a doughnut type is not in those 5 types, it cannot put into the doughnut box.

In test file for each way, it has one happy path and two error scenarios. About error ways, one is trying to put too
many doughnuts into a box, the other is an attempt to put something that isn’t doughnut into the box.

Two test ways, one is named table-driven test, the other one is named individual-subtest test. In Golang, table-driven
tests are the common way to write unit tests, individual-subtest is less common.

Table-driven tests offer a condensed way to write test scenarios while keeping code repetition to a minimum. The syntax
density is coming with the drawback of poor readability.

Individual-subtest syntax, unlike table-driven tests, favours readability overall “code dryness”. The tests may look
more repetitive, but each individual-subtest represents an individual “story”. As such, it is easier to follow
behaviour-driven testing principles with individual subtests than with table-driven tests.