1. open a terminal, % go run main.go

2. open the second terminal, try to find the main.go processID % ps -e

3. In the second terminal, make the main.go run in the background % kill -SIGHUP processID

check the first terminal, you will see the main.go has stopped, and it will show "[1] processID hangup go run main.go"

**4. Find current port-number whether occupy and return the port processID**
% lsof -i tcp:port-number

5. Kill the processID bound with port-number % kill processID // in the first terminal, it will show "[1] 81676
   terminated go run main.go"
   or % kill -SIGHUP processID     
   // it will show // "[1] processID hangup go run main.go"
   // "2022/12/30 23:02:43 hot reload"

6. Kill the port-number processID forever % kill -9 processID
