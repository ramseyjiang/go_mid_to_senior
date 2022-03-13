package singleton

import (
	"testing"
)

// go test -v -run TestGetInstance
func TestGetInstance(t *testing.T) {
	s := getInstance()
	if s == nil {
		t.Fatalf("First sigletone is nil")
	}
	s.SetTitle("First value")
	checkTitle := s.GetTitle()
	if checkTitle != "First value" {
		t.Errorf("First value is not setted")
	}

	s2 := getInstance()
	if s2 != s {
		t.Error("New instance different")
	}
	s2.SetTitle("New title")
	newTitle := s.GetTitle()
	if newTitle != "New title" {
		t.Errorf("Title different after change")
	}
}

func Test_singleton_GetTitle(t *testing.T) {
	type fields struct {
		title string
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &singleton{
				title: tt.fields.title,
			}
			if got := s.GetTitle(); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleton_SetTitle(t *testing.T) {
	type fields struct {
		title string
	}
	type args struct {
		t string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &singleton{
				title: tt.fields.title,
			}
			s.SetTitle(tt.args.t)
		})
	}
}
