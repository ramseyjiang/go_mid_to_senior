package sort

// quickSort Time complexity is O(nlogn), the best is O(nlogn), the worst is O(n*n)
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1

	for head < tail {
		if arr[i] > mid {
			// This is for go special change number.
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}

	arr[head] = mid
	quickSort(arr[:head])
	quickSort(arr[head+1:])

	return arr
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	num := length / 2
	left := mergeSort(arr[:num])
	right := mergeSort(arr[num:])

	return merge(left, right)
}

// merge Time complexity is O(nlogn), the best and the worst are the same, it is O(nlogn)
func merge(left, right []int) (result []int) {
	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			result = append(result, left[lIndex])
			lIndex++
		} else {
			result = append(result, right[rIndex])
			rIndex++
		}
	}

	result = append(result, left[lIndex:]...)
	result = append(result, right[rIndex:]...)
	return result
}

// bubbleSort Time complexity is O(n*n)
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

// Time complexity is O(n*n), the best is O(n), the worst is O(n*n)
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// selectionQuick Time complexity is O(n*n), the best is O(n*n), the worst is O(n*n)
func selectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	min := 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			// This is for go special change number.
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}
