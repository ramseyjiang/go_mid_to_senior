package ratelimit

var uniqueClientReqs = make(map[int]int)

func IsAllow(clientID int) (res bool) {
	limit := 5
	uniqueClientReqs[clientID] += 1
	if uniqueClientReqs[clientID] <= limit {
		return true
	}

	return false
}
