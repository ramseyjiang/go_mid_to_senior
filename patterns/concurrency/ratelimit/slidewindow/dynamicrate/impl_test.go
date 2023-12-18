package dynamicrate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name             string
		mockCurrentLoad  func() int
		numberOfRequests int
		expectedStatus   int
	}{
		{"LowLoad", func() int { return 30 }, 80, http.StatusOK},
		{"HighLoad", func() int { return 60 }, 60, http.StatusTooManyRequests},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := NewRequestWindow()
			handler := RateLimitMiddleware(rw, tt.mockCurrentLoad, http.HandlerFunc(RequestHandler))

			for i := 0; i < tt.numberOfRequests; i++ {
				req := httptest.NewRequest("GET", "/service", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if i == tt.numberOfRequests-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("%s: handler returned wrong status code: got %v want %v",
							tt.name, status, tt.expectedStatus)
					}
				}
			}
		})
	}
}
