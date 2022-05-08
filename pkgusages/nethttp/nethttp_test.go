package nethttppkg

import (
	"net/http"
	"testing"
)

func TestTrigger(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trigger()
		})
	}
}

func TestBaseUsage(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseUsage()
		})
	}
}

func TestHandler(t *testing.T) {
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.writer, tt.args.request)
		})
	}
}
