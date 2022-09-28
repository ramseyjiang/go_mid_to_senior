package main

func main() {
	_ = stackIt2()
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}
