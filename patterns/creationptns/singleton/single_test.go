package singleton

import (
	"testing"
)

// go test -v -run TestGetInstance
func TestGetInstance(t *testing.T) {
	s := GetInstance()
	if s == nil {
		t.Fatalf("First sigletone is nil")
	}
	s.SetTitle("First value")
	checkTitle := s.GetTitle()
	if checkTitle != "First value" {
		t.Errorf("First value is not setted")
	} else {
		t.Log("GetTitle and SetTitle work very well.")
	}

	s2 := GetInstance()
	if s2 != s {
		t.Error("New instance different")
	} else {
		t.Log("s2 equal with s, because they are all using Singleton struct.")
	}
	s2.SetTitle("New title")
	newTitle := s.GetTitle()
	if newTitle != "New title" {
		t.Errorf("Title different after change")
	} else {
		t.Log("GetTitle and SetTitle work very well.")
	}
}
