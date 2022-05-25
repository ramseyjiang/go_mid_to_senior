package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

type Status struct {
	Status string
}

type Test struct {
	description string
}

var ITEMS = 10000
var testArray [10000]Test

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UP")
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	var status = Status{
		"OK",
	}
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		return
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetHomePage)
	router.HandleFunc("/status", GetStatus)
	// router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	// register router manually
	router.HandleFunc("/debug/pprof/", pprof.Index)

	// The command line invocation of the current program
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)

	// CPU profile. You can specify the duration in the seconds GET parameter.
	// After you get the profile file, use the go tool pprof command to investigate the profile.
	router.HandleFunc("/debug/pprof/cpu", pprof.Profile)

	// Symbol looks up the program counters listed in the request, responding with a table mapping program counters to function names.
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// A trace of execution of the current program. You can specify the duration in the seconds GET parameter.
	// After you get the trace file, use the go tool trace command to investigate the trace.
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// A sampling of all past memory allocations
	router.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))

	// stack traces of holders of contended mutexes
	router.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

	// stack traces of all current goroutines
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))

	// A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))

	// stack traces that led to the creation of new OS threads
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// stack traces that led to blocking on synchronization primitives
	router.Handle("/debug/pprof/block", pprof.Handler("block"))

	log.Fatal(http.ListenAndServe(":1234", router))
}

func populateArray() {
	for i := 0; i < ITEMS; i++ {
		testArray[i] = Test{description: "Test " + string(i)}
	}
}

func main() {
	populateArray()
	handleRequests()
}
