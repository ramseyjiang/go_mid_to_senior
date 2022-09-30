package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ { // Adds 1 million elements
		m[i] = [128]byte{}
	}
	printAlloc()

	for i := 0; i < n; i++ { // Deletes 1 million elements
		delete(m, i)
	}

	runtime.GC() // Triggers a manual GC
	printAlloc()
	runtime.KeepAlive(m) // Keeps a reference to m so that the map isn’t collected
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m) // populates m with memory allocator statistics.
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
