package minwindowsubstr

func minSubstrWindow(s string, t string) string {
	m, n := len(s), len(t)
	if n == 0 || m < n {
		return ""
	}

	// Step 1: Create a frequency map for characters in t
	targetFreq := make(map[byte]int)
	for i := 0; i < n; i++ {
		targetFreq[t[i]]++
	}

	// Step 2: Sliding window setup
	left, right := 0, 0
	minStart, matchedCount := 0, 0
	minLen := m + 1
	windowFreq := make(map[byte]int)

	// Step 3: Expand the window by moving the `right` pointer
	for right < m {
		char := s[right]
		windowFreq[char]++

		// If char in window meets target frequency exactly, count it as matched
		if targetFreq[char] > 0 && windowFreq[char] == targetFreq[char] {
			matchedCount++
		}

		// Step 4: Shrink the window by moving `left` pointer
		for matchedCount == len(targetFreq) {
			// Check if the current window is smaller than previously found window
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			// Remove s[left] from the window and update counts
			leftChar := s[left]
			if targetFreq[leftChar] > 0 && windowFreq[leftChar] == targetFreq[leftChar] {
				matchedCount--
			}
			windowFreq[leftChar]--
			left++
		}

		right++
	}

	// Step 5: Return result
	if minLen == m+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
