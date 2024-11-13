package containermostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	width, h, area, maxAreaValue := 0, 0, 0, 0

	for left < right {
		width = right - left

		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		area = width * h
		maxAreaValue = max(maxAreaValue, area)
	}

	return maxAreaValue
}
