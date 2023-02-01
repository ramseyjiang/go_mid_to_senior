package folders

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Node{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Node{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")

	if cloneFolder == folder2 {
		t.Error("cloneFolder cannot be equal to folder2")
	}
}
