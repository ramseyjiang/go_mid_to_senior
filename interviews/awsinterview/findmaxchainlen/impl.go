package findmaxchainlen

// findMaxChainLength uses an array (size 26) to track the last occurrence of each character.
// For each starting index i, find the farthest valid ending index j by checking all characters larger than s[i].
// Directly access the array to get the farthest valid j.
// Time complexity is O(n). Space complexity is a map for 26 characters → O(1).
// Data structure is array, it is direct index access, so the look-up efficiency is faster direct addressing.
func findMaxChainLength(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	// Create lastOccurrence array uses to record every character in the string the last indices.
	// Characters in the string have the last indices and the position of the string.
	// Others are default value 0.
	lastOccurrence := make([]int, 26)
	for i, c := range s {
		// c's type is rune, it is int32. Using byte(c) to convert it to uint8
		// byte(c) ensures efficient character arithmetic for English letters (‘a’ - ‘z’).
		// Because English lowercase letters ('a' to 'z') fit within a byte (uint8) range (0-255).
		idx := byte(c) - 'a'
		lastOccurrence[idx] = i
	}

	maxLen := 0
	for i := 0; i < n; i++ {
		currentChar := s[i]
		maxJ := 0

		// iterate all characters' ASCII bigger than currentChar,
		// check them whether their positions are bigger than the currentChar position.
		for d := currentChar + 1; d <= 'z'; d++ {
			idx := d - 'a'
			j := lastOccurrence[idx]

			// record j to maxJ, the maxJ will be the max indices suitable for all conditions.
			if j > i && j > maxJ {
				maxJ = j
			}
		}

		if maxJ != 0 {
			// Calculate the currentLen, if currentLen is gather than maxLen, using the currentLen replace the maxLen
			currentLen := maxJ - i + 1
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	if maxLen >= 2 {
		return maxLen
	}
	return 0
}

// Optimized function using rightmost character positions
// Data structure is hash map, so the look-up efficiency is slightly slower (hash collisions) than the array.
func findMaxChainLength1(s string) int {
	// Edge cases handling. If s has less than 2 characters, return 0 immediately.
	n := len(s)
	if n < 2 {
		return 0
	}

	rightmost := make(map[byte]int) // Stores rightmost position of each character
	maxLength := 0

	// Store the rightmost position of each character in the string
	for i := 0; i < n; i++ {
		rightmost[s[i]] = i
	}

	// Iterate to find the longest valid substring
	for start := 0; start < n-1; start++ {
		// using byte(s[start]) to convert the character at s[start] from a rune (int32) to a byte (uint8)
		// byte(s[start]) ensures efficient character arithmetic for English letters (‘a’ - ‘z’).
		// Because English lowercase letters ('a' to 'z') fit within a byte (uint8) range (0-255).
		for ch := byte(s[start]) + 1; ch <= 'z'; ch++ { // Only check larger characters
			if end, exists := rightmost[ch]; exists && end > start {
				maxLength = max(maxLength, end-start+1)
			}
		}
	}

	return maxLength
}

// findMaxChainLength finds the longest valid substring length
// Time complexity Worst case: O(n^2) if every pair needs to be checked.
// Best case: O(n) if early exits happen often.
func findMaxChainLength2(s string) int {
	// Edge cases handling. If s has less than 2 characters, return 0 immediately.
	n := len(s)
	if n < 2 {
		return 0
	}

	maxLength := 0

	// Sliding window approach.
	// The start pointer iterates over the string.
	// The end pointer expands the window until a valid condition is met (s[start] < s[end]).
	for start := 0; start < n-1; start++ {
		for end := start + 1; end < n; end++ {
			if s[start] < s[end] { // Valid substring condition
				// Efficiently Tracking the Longest Valid Substring
				maxLength = max(maxLength, end-start+1)
			}
		}
	}

	return maxLength
}
