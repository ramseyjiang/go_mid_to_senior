package filterrooms

import (
	"reflect"
	"testing"
)

func TestFilterRooms(t *testing.T) {
	tests := []struct {
		name           string
		treasureRooms  []string
		instructions   [][]string
		expectedResult []string
	}{
		{
			name:          "Test 1: treasure_rooms_1, instructions_1",
			treasureRooms: []string{"lily", "tulip", "violet", "rose"},
			instructions: [][]string{
				{"jasmin", "tulip"},
				{"lily", "tulip"},
				{"tulip", "tulip"},
				{"rose", "rose"},
				{"violet", "rose"},
				{"sunflower", "violet"},
				{"daisy", "violet"},
				{"iris", "violet"},
			},
			// "tulip": incoming from "jasmin" and "lily" (ignoring self-loop)
			// "violet": incoming from "sunflower", "daisy", "iris"
			expectedResult: []string{"tulip", "violet"},
		},
		{
			name:          "Test 2: treasure_rooms_2, instructions_1",
			treasureRooms: []string{"lily", "jasmin", "violet"},
			instructions: [][]string{
				{"jasmin", "tulip"},
				{"lily", "tulip"},
				{"tulip", "tulip"},
				{"rose", "rose"},
				{"violet", "rose"},
				{"sunflower", "violet"},
				{"daisy", "violet"},
				{"iris", "violet"},
			},
			// Neither "tulip" (its destination "tulip" is not a treasure) nor "violet" (its destination "rose" is not a treasure) qualifies.
			expectedResult: []string{},
		},
		{
			name:          "Test 3: treasure_rooms_3, instructions_2",
			treasureRooms: []string{"violet"},
			instructions: [][]string{
				{"jasmin", "tulip"},
				{"lily", "tulip"},
				{"tulip", "violet"},
				{"violet", "violet"},
			},
			// "tulip": incoming from "jasmin" and "lily", and its instruction points to "violet" (a treasure)
			expectedResult: []string{"tulip"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterRooms(tt.treasureRooms, tt.instructions)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
