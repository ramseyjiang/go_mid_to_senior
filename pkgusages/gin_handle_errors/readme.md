This pkg is customised a gin middleware for error handlers. It is a way to customised errors in gin.

1. Run the server first go run main.go

2. send request to the server curl --location --request GET 'http://localhost:8080'

3. Check the server, you will see the below:
   [GIN] 2023/01/09 - 21:26:50 | 500 | 111.042µs | 127.0.0.1 | GET      "/"
   Error #01: What is this
