package ginpkg

import (
	"testing"

	"github.com/gin-gonic/gin"
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

func Test_getResponseString(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getResponseString(tt.args.router)
		})
	}
}

func Test_jsonResponse(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonResponse(tt.args.router)
		})
	}
}

func Test_jsonResponseBind(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonResponseBind(tt.args.router)
		})
	}
}

func Test_jsonResponseBindURI(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonResponseBindURI(tt.args.router)
		})
	}
}

func Test_jsonResponseFullPath(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonResponseFullPath(tt.args.router)
		})
	}
}

func Test_jsonResponseParamAndQuery(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonResponseParamAndQuery(tt.args.router)
		})
	}
}

func Test_multiParamsFormResponse(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiParamsFormResponse(tt.args.router)
		})
	}
}

func Test_postJSONResponse(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postJSONResponse(tt.args.router)
		})
	}
}

func Test_queryAndFormPostResponse(t *testing.T) {
	type args struct {
		router *gin.Engine
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryAndFormPostResponse(tt.args.router)
		})
	}
}
