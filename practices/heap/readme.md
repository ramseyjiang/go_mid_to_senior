in the heap folder in cmd,

/*when the heap.go return has "&", it will have the "moved to heap", if not, it won't have that.*/ package main

func main() {
_ = stackIt2()
}

//go:noinline func stackIt2() *int { y := 2 res := y * 2 return &res } % go build -gcflags '-m -l'

# github.com/ramseyjiang/go_mid_to_senior/practices/heap

./heap1.go:10:2: moved to heap: res

Using the copy way, it will save more execution time than using the pointer way. Because using the pointer way, it will
use more time to do the garbage collection. So optimisation Golang, one of choice is to reduce using pointers, but how
to balance it that should analysis using tools. +------------+------+---------+ | | Copy | Pointer |
+------------+------+---------+ | Goroutines | 41 | 406965 | | Heap | 10 | 197549 | | Threads | 15 | 12943 | | bgsweep |
0 | 193094 | | STW | 0 | 397 |