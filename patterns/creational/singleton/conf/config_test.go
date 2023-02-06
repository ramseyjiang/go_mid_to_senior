package conf

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConfig(t *testing.T) {
	conf := GetInstance()
	if conf == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	conf.Set("app_name", "My App")
	conf2 := GetInstance()
	expectedConf := conf
	if conf2 != expectedConf {
		t.Error("Expected same instance in conf2 but it got a different instance")
	}
	conf2.Set("app_version", "1.0.0")

	assert.Equal(t, "My App", conf.Get("app_name"))
	assert.Equal(t, "1.0.0", conf2.Get("app_version"))
}
