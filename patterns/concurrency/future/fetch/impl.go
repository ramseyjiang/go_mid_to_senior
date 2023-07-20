package fetch

import (
	"io"
	"net/http"
)

type Result struct {
	Body []byte
	Err  error
}

type Future struct {
	result chan Result
}

func NewFuture(f func() (Result, error)) *Future {
	res := make(chan Result, 1)
	go func() {
		r, err := f()
		res <- Result{Body: r.Body, Err: err}
		close(res)
	}()
	return &Future{result: res}
}

func (f *Future) Get() Result {
	return <-f.result
}

func fetch(url string) (Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return Result{Body: body}, err
}
