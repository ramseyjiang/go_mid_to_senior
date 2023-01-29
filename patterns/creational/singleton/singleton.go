package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

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
// If commented line 26, 27, 28, 30, it also works. Here just one example for how to use once.Do().
func GetInstance() *singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = new(singleton) // or instance = &singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
