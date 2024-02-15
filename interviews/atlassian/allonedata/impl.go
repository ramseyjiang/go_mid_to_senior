package allonedata

import (
	"sort"
)

type AllOne struct {
	Data map[string]int
}

func Constructor() AllOne {
	return AllOne{
		Data: make(map[string]int),
	}
}

func (a *AllOne) Inc(key string) {
	a.Data[key]++
}

func (a *AllOne) Dec(key string) {
	if a.Data[key] > 0 {
		a.Data[key]--
	} else {
		delete(a.Data, key) // Remove the key if its count is zero
	}
}

func (a *AllOne) GetMaxKey() string {
	maxValue := 0
	maxStr := []string{}
	for k, v := range a.Data {
		if v > maxValue {
			maxValue = v
			maxStr = []string{k}
		}

		if v == maxValue {
			maxStr = append(maxStr, k)
		}
	}

	sort.Strings(maxStr) // Sort to ensure consistent output
	if len(maxStr) > 0 {
		return maxStr[len(maxStr)-1] // Return the lexicographically largest string
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	minStr := []string{}
	minValue := 0
	for k, v := range a.Data {
		minValue = v
		minStr = []string{k}
		if v < minValue {
			minValue = v
			minStr = []string{k}
		}

		if v == minValue {
			minStr = append(minStr, k)
		}
	}

	sort.Strings(minStr) // Sort to ensure consistent output
	if len(minStr) > 0 {
		return minStr[0] // Return the lexicographically smallest string
	}
	return ""
}
