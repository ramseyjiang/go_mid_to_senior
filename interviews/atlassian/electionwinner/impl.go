package electionwinner

import (
	"slices"
)

func winner(votes []string) (winner string, max int) {
	counts := make(map[string]int)

	for _, v := range votes {
		counts[v]++
		if counts[v] > max {
			max = counts[v]
			winner = v
		}

		if counts[v] == max && len(v) < len(winner) {
			winner = v
		}
	}

	return
}

// winner finds the election winner from the slice of votes
// it returns the winning candidate's name and the number of votes received.
func winner2(votes []string) (winnerStr string, maxVotes int) {
	voteCounts := make(map[string]int)
	var candidatesMaxVotes []string

	// Count the votes for each candidate
	for _, vote := range votes {
		voteCounts[vote]++
		if voteCounts[vote] > maxVotes {
			maxVotes = voteCounts[vote]
			// Reset the slice as we have a new candidate with more votes
			candidatesMaxVotes = []string{vote} // Here, cannot use append to replace, because it is only one winner.
		} else if voteCounts[vote] == maxVotes {
			// Append candidate to the slice for further tie-breaking
			candidatesMaxVotes = append(candidatesMaxVotes, vote)
		}
	}

	// Sort the candidates who have the maximum votes to find the lexicographically smallest one
	slices.Sort(candidatesMaxVotes)
	winnerStr = candidatesMaxVotes[0] // The first candidate after sorting is the lexicographically smallest

	return winnerStr, maxVotes
}
