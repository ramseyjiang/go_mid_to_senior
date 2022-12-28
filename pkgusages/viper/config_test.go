package viperpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -run TestEntry
func TestEntry(t *testing.T) {
	want := Entry()
	var expected Config
	expected.AppName = "awesome web"
	expected.LogLevel = "DEBUG"

	assert.Equal(t, want.AppName, "awesome web")
	assert.Equal(t, want.LogLevel, "DEBUG")
}
