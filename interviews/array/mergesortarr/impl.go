package mergesortarr

import (
	"slices"
)

func mergeSortedArr1(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// Start from the last index of both arrays
	i, j, k := m-1, n-1, m+n-1

	// Traverse both arrays from the end
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are any remaining elements in nums2
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func mergeSortedArr2(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func mergeSortedArr(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// the append automatically removes the zeros if they are beyond the m elements,
	// as it only takes the non-zero elements from nums1[:m]
	nums1 = append(nums1[:m], nums2...)
	slices.Sort(nums1)
}
