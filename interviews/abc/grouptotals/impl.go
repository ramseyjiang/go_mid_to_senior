package grouptotals

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GroupTotals(strArr []string) string {
	// Create a map to store the totals
	totals := make(map[string]int)

	// Process each key-value pair
	for _, pair := range strArr {
		// Split the string into key and value
		parts := strings.Split(pair, ":")
		key := parts[0]
		value, _ := strconv.Atoi(parts[1]) // Convert the value to an integer

		// Add the value to the same key
		totals[key] += value
	}

	// Create a slice for the keys to sort
	keys := make([]string, 0, len(totals))
	for key, value := range totals {
		// If the value is 0, it won't include.
		if value != 0 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	// Build the result string in-place
	// var builder strings.Builder
	// for i, key := range keys {
	// 	if i > 0 {
	// 		builder.WriteString(",")
	// 	}
	// 	builder.WriteString(fmt.Sprintf("%s:%d", key, totals[key]))
	// }
	//
	// return builder.String()

	// Build the result string
	var result []string
	for _, key := range keys {
		result = append(result, fmt.Sprintf("%s:%d", key, totals[key]))
	}

	return strings.Join(result, ",")
}
