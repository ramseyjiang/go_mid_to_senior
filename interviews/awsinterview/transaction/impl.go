package transaction

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// After optimise
func processLogsOptimized(logs []string, threshold int32) []string {
	userCounts := make(map[string]int32)

	for _, log := range logs {
		space1 := strings.Index(log, " ")
		space2 := strings.LastIndex(log, " ")

		sender := log[:space1]
		recipient := log[space1+1 : space2]

		userCounts[sender]++
		if sender != recipient {
			userCounts[recipient]++
		}
	}

	var overThreshold []string
	for user, count := range userCounts {
		if count >= threshold {
			overThreshold = append(overThreshold, user)
		}
	}

	sort.Slice(overThreshold, func(i, j int) bool {
		userI, _ := strconv.Atoi(overThreshold[i])
		userJ, _ := strconv.Atoi(overThreshold[j])
		return userI < userJ
	})

	return overThreshold
}

// Before optimise
func processLogsInitial(logs []string, threshold int32) []string {
	// A map to track the transaction count of each user.
	userCounts := make(map[string]int32)

	for _, log := range logs {
		parts := strings.Split(log, " ")
		sender, recipient := parts[0], parts[1]

		// Increment the transaction count for the sender.
		userCounts[sender]++

		// If the sender and recipient are not the same, increment for the recipient.
		if sender != recipient {
			userCounts[recipient]++
		}
	}

	// Filter out users who have transaction counts greater than or equal to the threshold.
	var usersOverThreshold []string
	for user, count := range userCounts {
		if count >= threshold {
			usersOverThreshold = append(usersOverThreshold, user)
		}
	}

	// Sort user IDs in numeric order.
	sort.Slice(usersOverThreshold, func(i, j int) bool {
		a, _ := strconv.Atoi(usersOverThreshold[i])
		b, _ := strconv.Atoi(usersOverThreshold[j])
		return a < b
	})

	return usersOverThreshold
}

func main() {
	logs := []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"}
	threshold := int32(2)
	result := processLogsOptimized(logs, threshold)
	fmt.Println(result) // Expected output: ["88", "99"]
}
