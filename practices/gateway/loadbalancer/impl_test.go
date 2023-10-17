package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Custom ResponseRecorder that implements http.CloseNotifier
type ResponseRecorder struct {
	*httptest.ResponseRecorder
	closeNotifyChan chan bool
}

func (rr *ResponseRecorder) CloseNotify() <-chan bool {
	return rr.closeNotifyChan
}

func NewResponseRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func TestLoadBalancer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock servers
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Server 1"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Server 2"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}))
	defer server2.Close()

	// Update the servers slice in balancer.go to use our mock servers
	servers = []string{
		server1.URL,
		server2.URL,
	}

	router := CreateRouter()

	t.Run("TestServerRotation", func(t *testing.T) {
		responses := make(map[string]int)

		for i := 0; i < 10; i++ {
			w := NewResponseRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			router.ServeHTTP(w, req)

			responses[w.Body.String()]++
		}

		assert.Equal(t, 5, responses["Server 1"])
		assert.Equal(t, 5, responses["Server 2"])
	})
}
