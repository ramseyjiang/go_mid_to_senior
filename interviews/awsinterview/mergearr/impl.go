package mergearr

import "sort"

func merge(intervals [][]int) (res [][]int) {
	if len(intervals) == 0 {
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]

		// Use the last element in the res slice. If only 1 element, it will the last one.
		// After the new elements in, it is also the last one in the res slice.
		last := res[len(res)-1]

		if current[0] <= last[1] {
			res[len(res)-1][1] = max(last[1], current[1])
		} else {
			res = append(res, current)
		}
	}

	return res
}
