package intersectiontwoarr

import "sort"

// intersect is using the countMap way to solve.
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	countMap := make(map[int]int)
	var result []int

	for _, num := range nums1 {
		countMap[num]++
	}

	for _, num := range nums2 {
		if count, ok := countMap[num]; ok && count > 0 {
			countMap[num]--
			result = append(result, num)
		}
	}

	return result
}

// intersect2 is used the quick sort way, it uses two pointers to find the common elements.
// Time complexity is O(n+m), the space complexity is O(1).
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	var result []int

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}

	return result
}
