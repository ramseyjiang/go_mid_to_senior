package mycalendar

type MyCalendar struct {
	events [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{events: make([][2]int, 0)}
}

func (m *MyCalendar) Book(startTime int, endTime int) bool {
	for _, event := range m.events {
		s, e := event[0], event[1]
		if startTime < e && endTime > s {
			return false
		}
	}
	m.events = append(m.events, [2]int{startTime, endTime})

	return true
}
