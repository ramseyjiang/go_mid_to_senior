1. go run log.go
2. go run user.go
3. go run gateway.go
4. curl --location --request POST 'http://localhost:8082/log/add'
5. curl --location --request POST 'http://localhost:8082/user/create' \
   --header 'Content-Type: application/json' \
   --data-raw '{"name": "Ramsey"}'
6. curl --location --request GET 'http://localhost:8082/user/list'     
   [{"id":"4d34131d-c587-43e9-b76a-3bfd01c883a2","name":"Ramsey"}]

In the terminal of "go run user.go", it will show the following Log message: {"action":"User added","user":{"id":1,"name":"Ramsey"}}
