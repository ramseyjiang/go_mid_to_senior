package puzzles

import "testing"

func TestTriggerPointer(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerPointer()
		})
	}
}

func Test_second(t *testing.T) {
	type args struct {
		sourceArray []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			second(tt.args.sourceArray)
		})
	}
}

func Test_third(t *testing.T) {
	type args struct {
		sourceArray []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			third(tt.args.sourceArray)
		})
	}
}
