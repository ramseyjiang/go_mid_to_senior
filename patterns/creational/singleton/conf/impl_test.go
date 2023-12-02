package conf

import (
	"testing"
)

func TestConfig(t *testing.T) {
	testcases := []struct {
		name   string
		input1 string
		input2 string
		want   string
	}{
		{
			name:   "test 1",
			input1: "My App",
			input2: "app_name",
			want:   "My App",
		},
		{
			name:   "test 2",
			input1: "1.0.0",
			input2: "app_version",
			want:   "1.0.0",
		},
		{
			name: "test 3",
			want: "false",
		},
	}

	firstInstance := GetInstance()
	if firstInstance == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "test 3" {
				firstInstance.Set(tt.input2, tt.input1)
				if tt.want != firstInstance.Get(tt.input2) {
					t.Errorf("Expected GetInstance() Get value is %v, got %v", tt.want, firstInstance.Get(tt.input2))
				}
			} else {
				secondInstance := GetInstance()
				if firstInstance != secondInstance != (tt.want != tt.want) {
					t.Error("expected same instance but got a different instance")
				}
			}
		})
	}
}
