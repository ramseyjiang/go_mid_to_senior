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
	if _, exists := fs.paths[path]; exists || len(path) == 0 || path == "/" {
		return false
	}

	// Check if parent path exists, if it has several levels parents, the parent will be a slice.
	parent := path[:len(path)-1]
	for i := len(parent) - 1; i >= 0; i-- {
		if parent[i] == '/' {
			parent = parent[:i]
			break
		}
	}

	if parent != "" && fs.paths[parent] == 0 {
		return false
	}

	fs.paths[path] = value
	return true
}

func (fs *FileSystem) Get(path string) int {
	if value, exists := fs.paths[path]; exists {
		return value
	}
	return -1
}
