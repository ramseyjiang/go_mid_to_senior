package folders

import "fmt"

// Folder is Concrete prototype also, and it implements print() and clone()
type Folder struct {
	children []Node
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Node {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Node
	for _, i := range f.children {
		copyOperation := i.clone()
		tempChildren = append(tempChildren, copyOperation)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
