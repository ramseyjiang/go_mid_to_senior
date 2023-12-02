package plus

import (
	"fmt"
	"sync"
)

var once sync.Once

type SingletonIFace interface {
	AddOne() int
}

type Singleton struct {
	count int
}

var instance *Singleton

// GetInstance is used to return a singleton instance.
func GetInstance() *Singleton {
	once.Do(
		func() {
			fmt.Println("Creating single instance now.")
			instance = new(Singleton) // or instance = &Singleton{}
		})
	return instance
}

func (s *Singleton) AddOne() int {
	s.count++
	return s.count
}
