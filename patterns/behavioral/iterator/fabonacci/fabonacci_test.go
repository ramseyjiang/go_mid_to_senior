package fabonacci

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFab(t *testing.T) {
	sum := 0
	f := &Fibonacci{a: 0, b: 1, max: 10}
	it := f.New()
	expected := 88
	for it.HasNext() {
		sum += it.Next()
	}

	assert.Equal(t, expected, sum)
}
