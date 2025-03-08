package validation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRegistration(t *testing.T) {
	tests := []struct {
		name          string
		payload       interface{}
		expectedCode  int
		expectError   bool
		errorContains []string
	}{
		{
			name: "Valid registration",
			payload: map[string]interface{}{
				"username": "valid_user123",
				"email":    "test@example.com",
				"password": "securepassword",
				"age":      25,
			},
			expectedCode: http.StatusOK,
			expectError:  false,
		},
		{
			name: "Short username",
			payload: map[string]interface{}{
				"username": "ab",
				"email":    "test@example.com",
				"password": "password",
				"age":      20,
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Username", "min"},
		},
		{
			name: "Invalid email format",
			payload: map[string]interface{}{
				"username": "validuser",
				"email":    "invalid-email",
				"password": "password123",
				"age":      30,
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Email", "email"},
		},
		{
			name: "Age under 18",
			payload: map[string]interface{}{
				"username": "young_user",
				"email":    "young@example.com",
				"password": "youngpass",
				"age":      17,
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Age", "gte"},
		},
		{
			name: "Missing required fields",
			payload: map[string]interface{}{
				"username": "missing_fields",
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Email", "Password", "Age"},
		},
	}

	router := setupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectError {
				var resp Response
				json.Unmarshal(w.Body.Bytes(), &resp)

				assert.False(t, resp.Success)
				for _, keyword := range tt.errorContains {
					assert.Contains(t, resp.Error, keyword)
				}
			} else {
				var resp Response
				json.Unmarshal(w.Body.Bytes(), &resp)

				assert.True(t, resp.Success)
				assert.NotEmpty(t, resp.Data.(map[string]interface{})["user"])
			}
		})
	}
}

func TestHandleOrders(t *testing.T) {
	validOrder := map[string]interface{}{
		"customer_id": "cust_123",
		"items": []map[string]interface{}{
			{
				"product_id": "prod_1",
				"name":       "Product 1",
				"quantity":   2,
				"price":      19.99,
			},
		},
		"total_amount": 39.98,
		"status":       "pending",
	}

	tests := []struct {
		name          string
		payload       interface{}
		expectedCode  int
		expectError   bool
		errorContains []string
	}{
		{
			name:         "Valid order",
			payload:      validOrder,
			expectedCode: http.StatusOK,
			expectError:  false,
		},
		{
			name: "Invalid order status",
			payload: map[string]interface{}{
				"customer_id": "cust_123",
				"items": []map[string]interface{}{
					{
						"product_id": "prod_1",
						"name":       "Product 1",
						"quantity":   2,
						"price":      19.99,
					},
				},
				"total_amount": 39.98,
				"status":       "invalid_status",
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Status", "oneof"},
		},
		{
			name: "Exceed item quantity",
			payload: map[string]interface{}{
				"customer_id": "cust_123",
				"items": []map[string]interface{}{
					{
						"product_id": "prod_1",
						"name":       "Product 1",
						"quantity":   11,
						"price":      19.99,
					},
				},
				"total_amount": 219.89,
				"status":       "processing",
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Quantity", "lte"},
		},
		{
			name: "Missing customer ID",
			payload: map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"product_id": "prod_1",
						"name":       "Product 1",
						"quantity":   1,
						"price":      10.0,
					},
				},
				"total_amount": 10.0,
				"status":       "pending",
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"CustomerID", "required"},
		},
		{
			name: "Long notes field",
			payload: map[string]interface{}{
				"customer_id": "cust_123",
				"items": []map[string]interface{}{
					{
						"product_id": "prod_1",
						"name":       "Product 1",
						"quantity":   1,
						"price":      10.0,
					},
				},
				"total_amount": 10.0,
				"status":       "completed",
				"notes":        string(make([]byte, 501)),
			},
			expectedCode:  http.StatusBadRequest,
			expectError:   true,
			errorContains: []string{"Notes", "max"},
		},
	}

	router := setupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectError {
				var resp Response
				json.Unmarshal(w.Body.Bytes(), &resp)

				assert.False(t, resp.Success)
				for _, keyword := range tt.errorContains {
					assert.Contains(t, resp.Error, keyword)
				}
			} else {
				var resp Response
				json.Unmarshal(w.Body.Bytes(), &resp)

				assert.True(t, resp.Success)
				assert.NotNil(t, resp.Data.(map[string]interface{})["order"])
			}
		})
	}
}

func TestGetErrorMsg(t *testing.T) {
	tests := []struct {
		field    string
		tag      string
		expected string
	}{
		{"Username", "required", "This field is required"},
		{"Price", "gt", "Must be greater than 0"},
		{"Quantity", "lte", "Must not exceed 10 items"},
		{"Status", "oneof", "Must be one of: pending, processing, completed"},
		{"Notes", "max", "Must not exceed 500 characters"},
		{"Unknown", "unknown", "Invalid value"},
	}

	for _, tt := range tests {
		t.Run(tt.tag+"_"+tt.field, func(t *testing.T) {
			assert.Equal(t, tt.expected, getErrorMsg(tt.field, tt.tag))
		})
	}
}
