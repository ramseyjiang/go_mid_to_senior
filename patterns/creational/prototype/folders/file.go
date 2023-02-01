package folders

import "fmt"

// Node is Prototype interface
type Node interface {
	print(string)
	clone() Node
}

// File is Concrete prototype, and it implements print() and clone()
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Node {
	return &File{name: f.name + "_clone"}
}
