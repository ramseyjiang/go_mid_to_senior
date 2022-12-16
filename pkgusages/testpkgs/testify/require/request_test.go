package require

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSavePlayerInfo is used individual-subtest way to test happy path and error scenarios.
func TestSavePlayerInfo(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte(`{"message": "success"`))
		if err != nil {
			panic("cannot return http response")
		}
	}))
	defer testServer.Close()

	t.Run("happy path", func(t *testing.T) {
		s := PlayerInfo{Name: "Denis Rodman", Team: "Chicago Bulls", Position: "Forward"}
		err := savePlayerInfo(s, testServer.URL)
		require.NoError(t, err)
	})

	t.Run("sad path: invalid stats", func(t *testing.T) {
		s := PlayerInfo{Name: "Denis Green", Team: "Warriors"}
		err := savePlayerInfo(s, testServer.URL)
		require.Error(t, err)
		assert.Equal(t, "missing data", err.Error())
	})
}
