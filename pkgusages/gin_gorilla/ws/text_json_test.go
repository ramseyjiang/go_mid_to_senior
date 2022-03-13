package ws

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJSONAPI(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JSONAPI(tt.args.c)
		})
	}
}

func TestTextAPI(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TextAPI(tt.args.c)
		})
	}
}
