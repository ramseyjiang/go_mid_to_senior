package reqmock

import (
	"fmt"

	"github.com/asmcos/requests"
)

func Trigger() {
	req := requests.Requests()
	resp, _ := req.Get("https://baidu.com")
	fmt.Println(resp.R.StatusCode)             // 200
	fmt.Println(resp.R.Header["Content-Type"]) // [text/html;charset=utf-8]
}
