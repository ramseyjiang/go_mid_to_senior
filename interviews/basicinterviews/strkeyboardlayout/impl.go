package strkeyboardlayout

import "strings"

// spell generates instructions to spell a given string on a provided keyboard layout.
func spell(layout [][]string, s string) string {
	// Map to store character positions in the layout
	pos := make(map[string][2]int)

	// Build the position map from the layout
	for i, row := range layout {
		for j, char := range row {
			pos[char] = [2]int{i, j}
		}
	}

	var instructions []string
	currPos := [2]int{0, 0} // Start at top-left corner

	// Iterate through each character in the string
	for _, char := range s {
		targetPos, exists := pos[string(char)]
		if !exists {
			continue // Skip if the character is not in the layout
		}

		// Compute the vertical and horizontal movements
		dRow := targetPos[0] - currPos[0]
		dCol := targetPos[1] - currPos[1]

		// Move vertically first (UP/DOWN)
		if dRow < 0 {
			instructions = append(instructions, strings.Repeat("U", -dRow))
		} else {
			instructions = append(instructions, strings.Repeat("D", dRow))
		}

		// Move horizontally (LEFT/RIGHT)
		if dCol < 0 {
			instructions = append(instructions, strings.Repeat("L", -dCol))
		} else {
			instructions = append(instructions, strings.Repeat("R", dCol))
		}

		// Press the character
		instructions = append(instructions, "P")

		// Update current position
		currPos = targetPos
	}

	return strings.Join(instructions, "")
}
