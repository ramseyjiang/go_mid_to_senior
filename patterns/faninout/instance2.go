package faninout

import (
	"fmt"
	"time"
)

func EntryInstance2() {
	source := generate("data string")
	process := getProcessor()

	for i := 0; i < 10; i++ {
		// Interact with the processor instance with the postJob method
		// Inside the postJob method, only when there is a worker.
		process.postJob(source)
	}
}

// generate is evident returns a reception from a channel.
func generate(data string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- data
			time.Sleep(1000)
		}
	}()

	return channel
}

type Processor struct {
	jobChannel chan string
	done       chan *Worker
	workers    []*Worker
}
type Worker struct {
	name string
}

func (w *Worker) processJob(data string, done chan *Worker) {
	// Use the data and process the job
	// The worker does the processing in a separate goroutine and notifies the processor instance via the done channel.
	go func() {
		fmt.Println("Working on data ", data, w.name)
		time.Sleep(3000)
		done <- w
	}()
}

// getProcessor creates an instance of the Processor and start its processing.
func getProcessor() *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers:    make([]*Worker, 5),
		done:       make(chan *Worker),
	}
	for i := 0; i < 5; i++ {
		w := &Worker{name: fmt.Sprintf("<Worker - %d>", i)}
		p.workers[i] = w
	}
	p.startProcess()
	return p
}

// StartProcess method we have 2 selects.
// Get the messages send by the generator.
func (p *Processor) startProcess() {
	go func() {
		for {
			select {

			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					// Select the worker which is always the top worker of the worker slice inside the processor instance.
					p.workers = p.workers[1:]

					// It can block if there is no jobs. In those cases make sure to add backpressure handling.
					w.processJob(<-p.jobChannel, p.done)
				}
			case w := <-p.done: // The signal from the worker is caught and the worker is appended to the worker list.
				p.workers = append(p.workers, w)
			}
		}
	}()
}

// Get the message from the generator and channel it to the jobChannel channel in the processor instance.
func (p *Processor) postJob(jobs <-chan string) {
	p.jobChannel <- <-jobs
}
