package gorillaws

import (
	"net/http"
	"testing"
)

func TestTriggerServer(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TriggerServer()
		})
	}
}

func Test_home(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home(tt.args.w, tt.args.r)
		})
	}
}

func Test_socketHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			socketHandler(tt.args.w, tt.args.r)
		})
	}
}
