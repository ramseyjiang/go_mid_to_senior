package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance is used to return a singleton instance.
// If instance is not nil, it means the instance has been new as a singleton.
// If instance is nil, it means the instance has not new yet, so new it and return.
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
