package ratelimit

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

func get(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

// Task struct is used to store the result of each request and response.
type Task struct {
	taskNum    int
	url        string
	data       []byte
	err        string
	statusCode int
	duration   time.Duration
	workerName string
}

// worker is used to handle the task from the taskChan, and send the result to the resultsChan.
// taskChan is requests container chan,
// resultChan is responses container chan,
func worker(name string, taskChan <-chan Task, resultChan chan<- Task) {
	for task := range taskChan {
		start := time.Now()
		body, code, err := get(task.url)
		if err != nil {
			task.err = err.Error()
		}
		task.statusCode = code
		task.data = body
		task.workerName = name
		task.duration = time.Duration(time.Since(start).Milliseconds())
		resultChan <- task
	}
}

// producer component generates all tasks and sends them to Task Chan for the downstream process.
func producer(rl ratelimit.Limiter) <-chan Task {
	var tasks []Task
	for i := 0; i < 5; i++ {
		url := "https://httpbin.org/get?i=" + fmt.Sprintf("%d", i)
		tasks = append(tasks, Task{taskNum: i, url: url})
	}
	fmt.Println("total urls: ", len(tasks))

	out := make(chan Task)
	go func() {
		defer close(out)
		for _, task := range tasks {
			rl.Take()
			out <- task
		}
	}()

	return out
}
