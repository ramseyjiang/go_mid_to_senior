package setintersectionsize

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	count := 2
	first, second := intervals[0][1], intervals[0][1]-1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] < first || intervals[i][0] > first {
			// first is invalid, re-assign first
			// set second by first
			first, second = intervals[i][1], first
			count++
		}
		if intervals[i][1] < second || intervals[i][0] > second {
			// second is invalid, re-assign second
			second = intervals[i][1]
			count++
			if second == first {
				// same as first, set by first-1
				second = first - 1
			} else if second > first {
				// sort
				first, second = second, first
			}
		}
	}
	return count
}
