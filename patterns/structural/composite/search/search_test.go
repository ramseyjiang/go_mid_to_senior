package search

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSearch(t *testing.T) {
	file1 := &File{name: "File1"}
	assert.Equal(t, "File1", file1.getName())

	file2 := &File{name: "File2"}
	assert.Equal(t, "File2", file2.getName())

	file3 := &File{name: "File3"}
	assert.Equal(t, "File3", file3.getName())

	folder1 := &Folder{
		name: "Folder1",
	}
	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	assert.Equal(t, 3, len(folder2.components))

	folder2.search("rose")
}
