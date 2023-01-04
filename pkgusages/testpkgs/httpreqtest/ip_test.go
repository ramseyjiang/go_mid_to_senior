package httpreqtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	wanted := "122.60.180.180"
	expected := getIP()
	assert.Equal(t, expected, wanted)
}
