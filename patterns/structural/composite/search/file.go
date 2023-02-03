package search

import "fmt"

type Component interface {
	search(string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
