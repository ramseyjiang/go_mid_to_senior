package getmaxstablity

import (
	"sort"
)

// getMaxStability calculates the maximum stability of any subset of servers.
// Stability is defined as (minimum availability in subset) * (sum of reliabilities in subset).
// The result is returned modulo 10^9 + 7.
func getMaxStability(reliability []int32, availability []int32) int32 {
	const mod int32 = 1000000007 // Modulo value for the result
	n := len(reliability)

	// Create a slice of structs to pair availability and reliability for each server
	servers := make([]struct{ a, r int32 }, n)
	for i := 0; i < n; i++ {
		servers[i].a = availability[i]
		servers[i].r = reliability[i]
	}

	// Sort servers in descending order of availability.
	// This ensures that for the first i servers, the minimum availability is servers[i].a.
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].a > servers[j].a
	})

	var maxStability int64 = 0   // Track maximum stability (int64 to avoid overflow)
	var sumReliability int64 = 0 // Cumulative sum of reliabilities (int64 to avoid overflow)

	// Iterate through sorted servers and compute stability for subsets
	for i := 0; i < n; i++ {
		sumReliability += int64(servers[i].r)           // Add current server's reliability to the cumulative sum
		current := int64(servers[i].a) * sumReliability // Stability = min_availability * sum_reliability
		if current > maxStability {
			maxStability = current // Update maximum stability if current subset is better
		}
	}

	// Apply modulo and convert back to int32
	result := maxStability % int64(mod)
	return int32(result)
}
