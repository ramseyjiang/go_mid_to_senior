package gozero

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/service"
)

func morningStart(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	_, err := fmt.Fprintln(w, "morning!")
	if err != nil {
		return
	}
}

func eveningStart(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	_, err := fmt.Fprintln(w, "evening!")
	if err != nil {
		return
	}
}

type Morning struct{}

func (m Morning) Start() {
	http.HandleFunc("/morning", morningStart)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func (m Morning) Stop() {
	fmt.Println("Stop morning service...")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening", eveningStart)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		return
	}
}

func (e Evening) Stop() {
	fmt.Println("Stop evening service...")
}

// Trigger is used to run wayOne, wayTwo and wayThree, only one can be executed in each run time.
// If you want to run it, please replace trigger to main.
func Trigger() {
	wayOne()
	wayTwo()
	wayThree()
}

func wayOne() {
	fmt.Println("WayOne Start morning service...")
	var morning Morning
	morning.Start()
	defer morning.Stop()

	fmt.Println("WayOne Start evening service...")
	var evening Evening
	evening.Start()
	evening.Stop()
}

func wayTwo() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("WayTwo Start morning service...")
		var morning Morning
		defer morning.Stop()
		morning.Start()
	}()
	go func() {
		defer wg.Done()
		fmt.Println("WayTwo Start evening service...")
		var evening Evening
		defer evening.Stop()
		evening.Start()
	}()
	wg.Wait()
}

func wayThree() {
	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(Morning{})
	group.Add(Evening{})
	group.Start()
}
