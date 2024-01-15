package rankteamsvotes

import "sort"

func RankVotes(votes []string) string {
	if len(votes) == 0 {
		return ""
	}

	// Teams in a default order, needing sorting by round winners.
	// Convert the string votes[0] to byte
	teams := []byte(votes[0])

	// Capture round votes
	roundVotes := [26][26]int{}
	for _, vote := range votes {
		for i, team := range vote {
			// team-'A' means to get the team number, i means the team votes
			roundVotes[team-'A'][i]++
		}
	}

	// Determine winning team each round
	sort.SliceStable(teams, func(i, j int) bool {
		a := roundVotes[teams[i]-'A']
		b := roundVotes[teams[j]-'A']
		for round := 0; round < 26; round++ {
			if a[round] == b[round] {
				// Tied, so need to keep looking for a winner
				continue
			}
			// Found a round winner
			return a[round] > b[round]
		}
		// No winner, so just sort alphabetically
		return teams[i] < teams[j]
	})

	return string(teams)
}
