package onlineelection

import (
	"sort"
)

type TopVotedCandidate struct {
	persons []int
	times   []int
}

func NewOnlineElection(persons []int, times []int) TopVotedCandidate {
	counts := make([]int, len(persons))
	var maxCount int
	var maxPerson int
	for i, person := range persons {
		counts[person]++
		if counts[person] >= maxCount {
			maxCount = counts[person]
			maxPerson = person
		}
		persons[i] = maxPerson
	}

	return TopVotedCandidate{
		persons: persons,
		times:   times,
	}
}

func (tvc *TopVotedCandidate) Query(t int) int {
	i := sort.Search(len(tvc.times), func(i int) bool {
		return tvc.times[i] > t
	})
	return tvc.persons[i-1]
}
