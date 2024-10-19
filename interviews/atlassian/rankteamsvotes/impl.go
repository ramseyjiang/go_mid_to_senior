package rankteamsvotes

import "sort"

func RankVotes(votes []string) string {
	output := ""
	if len(votes) == 0 {
		return output
	}

	if len(votes) == 1 {
		return votes[0]
	}

	// Get all char counts.
	counts := make(map[rune]int)
	keys := make([]rune, 0)
	for _, vote := range votes {
		for k, v := range vote {
			if _, ok := counts[v]; ok {
				counts[v] += k
			} else {
				keys = append(keys, v) // Get all chars slice
				counts[v] += k
			}
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		// If counts are equal, sort alphabetically
		if counts[keys[i]] == counts[keys[j]] {
			return keys[i] < keys[j]
		}
		// Otherwise, sort by counts (ascending order)
		return counts[keys[i]] < counts[keys[j]]
	})

	for _, k := range keys {
		output += string(k)
	}

	return output
}

func RankVotes1(votes []string) string {
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

func RankVotes2(votes []string) string {
	if len(votes) == 0 {
		return ""
	}

	// Create a map to hold the count of each rank for each team.
	rankCount := make(map[rune][]int)
	for i := range votes[0] {
		rankCount[rune(votes[0][i])] = make([]int, len(votes[0]))
	}

	// Count the ranks for each team.
	for _, vote := range votes {
		for rank, team := range vote {
			rankCount[team][rank]++
		}
	}

	// Sort the teams according to the problem's rules.
	teams := []rune(votes[0])
	sort.Slice(teams, func(i, j int) bool {
		for k := range rankCount[teams[i]] {
			if rankCount[teams[i]][k] != rankCount[teams[j]][k] {
				return rankCount[teams[i]][k] > rankCount[teams[j]][k]
			}
		}
		return teams[i] < teams[j]
	})

	// Construct the final sorted string of teams.
	result := make([]byte, len(teams))
	for i, team := range teams {
		result[i] = byte(team)
	}
	return string(result)
}
