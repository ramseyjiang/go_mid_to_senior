package filesystem

type FileSystem struct {
	paths map[string]int
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		paths: make(map[string]int),
	}
}

func (fs *FileSystem) CreatePath(path string, value int) bool {
	// Already exists, empty path, or just '/'
	if _, exists := fs.paths[path]; exists || len(path) == 0 || path == "/" {
		return false
	}

	// Check if parent path exists, if it has several levels parents, the parent will be a slice.
	parent := path[:len(path)-1]
	for i := len(parent) - 1; i >= 0; i-- {
		if parent[i] == '/' {
			parent = parent[:i] // Extract parent path
			break
		}
	}

	if parent != "" && fs.paths[parent] == 0 {
		return false // Parent path doesn't exist
	}

	fs.paths[path] = value
	return true
}

func (fs *FileSystem) Get(path string) int {
	if value, exists := fs.paths[path]; exists {
		return value
	}
	return -1 // Path doesn't exist
}

/*
package filesystem

import (
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	value    int
	isEnd    bool
}

type FileSystem struct {
	root *TrieNode
}

func NewFileSystem() FileSystem {
	return FileSystem{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
	}
}

func (fs *FileSystem) CreatePath(path string, value int) bool {
	if path == "" || path == "/" || fs.root == nil {
		return false
	}

	parts := strings.Split(path, "/")[1:] // split and remove the first empty part
	node := fs.root

	for i, part := range parts {
		if next, exists := node.children[part]; exists {
			node = next
		} else {
			if i == len(parts)-1 {
				// Only create the final part if we're at the end of the path
				node.children[part] = &TrieNode{
					children: make(map[string]*TrieNode),
					value:    value,
					isEnd:    true,
				}
				return true
			}
			return false
		}
	}

	// If we exit the loop without creating a new node, the path already exists
	if node.isEnd {
		return false
	}

	node.value = value
	node.isEnd = true
	return true
}

func (fs *FileSystem) Get(path string) int {
	if path == "" || path == "/" || fs.root == nil {
		return -1
	}

	parts := strings.Split(path, "/")[1:] // split and remove the first empty part
	node := fs.root

	for _, part := range parts {
		if next, exists := node.children[part]; exists {
			node = next
		} else {
			return -1
		}
	}

	if node.isEnd {
		return node.value
	}
	return -1
}
*/
