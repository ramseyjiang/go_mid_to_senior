The OPA policy is a JSON document that defines the permissions of your users. OPA is an open source project that
provides a policy engine. It can be used to enforce policies in your application. It can be used to authorize your
users, validate your data, and much more. By using OPA, policies are global to the organization and changes doesn’t
force you to modify your code.

First step, % go run main.go

Second step, Open postman, and access the auth_opa collection.

Third step, Access the localhost:8080/articles, with the "X-User" in the request header.

Fourth step, When you change the "X-User" value, among "alice, bob and ramsey", you will see responses are different.
