1. open a terminal, % go run main.go

2. open the other teminal % ps -e

3. In the second terminal, try to find the main.go processID
4. In the second terminal % kill -SIGHUP processID

5. check the first terminal, you will see the main.go has stopped, and it will show "[1] processID hangup go run
   main.go"
