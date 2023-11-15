package reversestr

import (
	"reflect"
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "test 1",
			input: []byte{'h', 'e', 'l', 'l', 'o'},
			want:  []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			"test 2",
			[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
			[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseStr(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
