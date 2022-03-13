package practices

import (
	"reflect"
	"testing"

	"golang_learn/customizepkgs/goerr"
)

func TestResourceNotFound(t *testing.T) {
	type args struct {
		id      string
		message string
		kind    string
		cause   error
	}
	var tests []struct {
		name string
		args args
		want goerr.Error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceNotFound(tt.args.id, tt.args.message, tt.args.kind, tt.args.cause); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriggerCelsiusConvert(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerCelsiusConvert()
		})
	}
}

func TestTriggerErrorPkg(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerErrorPkg()
		})
	}
}

func TestTriggerGreet(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerGreet()
		})
	}
}

func TestTriggerKeyboardInput(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerKeyboardInput()
		})
	}
}
