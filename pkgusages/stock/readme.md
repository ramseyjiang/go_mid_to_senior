This is using echo framework to build. The stock info is from yahoo.com.

1. After the server run, access the "hello world"
   curl --location --request GET 'http://localhost:8001'

The logger is below:
{"time":"2023-01-09T20:50:27.875106+13:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8000","method":"GET","
uri":"/","user_agent":"curl/7.84.0","status":200,"error":"","latency":6792,"latency_human":"6.792µs","bytes_in":0,"
bytes_out":14}

2. Access the stock price is using post.

curl X 'http://localhost:8001/price' \
--header 'Content-Type: application/json' \
--data-raw '{
"ticker":"aapl"
}'

return is
"Apple Inc. (AAPL)NasdaqGS - NasdaqGS Real Time Price. Currency in USDFollowVisitors trend2W10W9M129.62+4.60 (+3.68%)At
close:  04:00PM EST129.25 -0.37 (-0.29%)After hours: 07:59PM EST \n Advertisement\n            "

The logger is:
"time":"2023-01-09T20:58:00.44663+13:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8000","method":"POST","uri":"
/price","user_agent":"curl/7.84.0","status":200,"error":"","latency":728156916,"latency_human":"728.156916ms","
bytes_in":19,"bytes_out":233}
