package majorityelement

func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for _, n := range nums {
		if n == majority {
			count++
		} else {
			count--
			if count == 0 {
				majority = n
				count = 1
			}
		}
	}

	// Verify if the found majority element appears more than N/2 times
	count = 0
	for _, n := range nums {
		if n == majority {
			count++
		}

		if count > len(nums)/2 {
			return majority
		}
	}

	// If no element appears more than N/2 times, return -1
	return -1
}

func majorityElement2(nums []int) int {
	freqMap := make(map[int]int)

	// Iterate through the array and store element frequencies in the map
	for _, num := range nums {
		freqMap[num]++
	}

	// Check if any element has a frequency greater than N/2
	for num, count := range freqMap {
		if count > len(nums)/2 {
			return num
		}
	}

	// No majority element found
	return -1
}
