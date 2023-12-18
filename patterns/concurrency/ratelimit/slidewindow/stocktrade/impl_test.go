package stocktrade

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTradeLimitMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		numberOfTrades int
		sleepDuration  time.Duration
		expectedStatus int
	}{
		{"UnderLimit", 4, 0, http.StatusOK},
		{"AtLimit", 5, 0, http.StatusOK},
		{"OverLimit", 6, 0, http.StatusTooManyRequests},
		{"ResetAfterOneMinute", 6, 1 * time.Minute, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new TradeWindow for each test to ensure isolation
			tw := NewTradeWindow()
			handler := TradeLimitMiddleware(tw, http.HandlerFunc(TradeHandler))

			for i := 0; i < tt.numberOfTrades; i++ {
				if i == 5 && tt.sleepDuration > 0 {
					time.Sleep(tt.sleepDuration)
				}

				req := httptest.NewRequest("POST", "/trade", nil)
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if i == tt.numberOfTrades-1 {
					if status := rr.Code; status != tt.expectedStatus {
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}
		})
	}
}
