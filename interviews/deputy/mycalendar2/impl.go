package mycalendar2

type MyCalendarTwo struct {
	booked   [][]int
	overlaps [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booked:   make([][]int, 0),
		overlaps: make([][]int, 0),
	}
}

func (m *MyCalendarTwo) Book(startTime int, endTime int) bool {
	// Until the overlaps are not empty, this loop will check.
	for _, interval := range m.overlaps {
		s, e := interval[0], interval[1]
		if startTime < e && endTime > s {
			return false
		}
	}

	newOverlaps := make([][]int, 0)
	// Until the booked is not empty, this loop will check.
	for _, interval := range m.booked {
		s, e := interval[0], interval[1]
		overlapStart := max(startTime, s)
		overlapEnd := min(endTime, e)

		// Only one condition add start and end to overlaps.
		if overlapStart < overlapEnd {
			newOverlaps = append(newOverlaps, []int{overlapStart, overlapEnd})
		}
	}

	m.overlaps = append(m.overlaps, newOverlaps...)
	m.booked = append(m.booked, []int{startTime, endTime})

	return true
}
