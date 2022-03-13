package singleton

import (
	"fmt"
	"sync"
)

// declare variable
var once sync.Once
var singleInstance *singleton

// declaration defined type
type singleton struct {
	title string
}

// Singleton defined type with interface
type Singleton interface {
	SetTitle(t string)
	GetTitle() string
}

// getInstance is used to get only one object.
func getInstance() Singleton {
	if singleInstance == nil {
		once.Do(func() { // once.Do is used to prevent multiple goroutines are trying to access one instance together.
			fmt.Println("Creating single instance now.")
			singleInstance = new(singleton)
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// SetTitle is a Setter for singleton variable
func (s *singleton) SetTitle(t string) {
	s.title = t
}

// GetTitle is a Getter singleton variable
func (s *singleton) GetTitle() string {
	return s.title
}
