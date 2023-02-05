package conf

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConfig(t *testing.T) {
	conf := GetInstance()
	conf.Set("app_name", "My App")
	conf.Set("app_version", "1.0.0")

	assert.Equal(t, "My App", conf.Get("app_name"))
	assert.Equal(t, "1.0.0", conf.Get("app_version"))
}
