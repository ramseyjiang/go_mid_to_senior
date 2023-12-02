package plus

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	firstInstance := GetInstance() // Initialize the singleton instance
	if firstInstance == nil {
		t.Fatal("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	tests := []struct {
		name string
		want int
	}{
		{
			name: "First Increment",
			want: 1,
		},
		{
			name: "Second Increment",
			want: 2,
		},
		{
			name: "Check Same Instance",
			want: 0,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Check Same Instance" {
				secondInstance := GetInstance()
				if (firstInstance != secondInstance) != (tt.want != 0) {
					t.Error("expected same instance but got a different instance")
				}
			} else {
				got := firstInstance.AddOne()
				if got != tt.want {
					t.Errorf("expected increment to be %v, got %v", tt.want, got)
				}
			}
		})
	}
}
