package dynamicrate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	mockGetCurrentLoad := func() int {
		return 30 // Mocked current load value
	}

	tests := []struct {
		name             string
		numberOfRequests int
		expectedStatus   int
	}{
		{"LowLoad", 80, http.StatusOK},
		{"HighLoad", 60, http.StatusTooManyRequests},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := NewRequestWindow()
			handler := RateLimitMiddleware(rw, mockGetCurrentLoad, http.HandlerFunc(RequestHandler))

			for i := 0; i < tt.numberOfRequests; i++ {
				req := httptest.NewRequest("GET", "/service", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if i == tt.numberOfRequests-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}
		})
	}
}
