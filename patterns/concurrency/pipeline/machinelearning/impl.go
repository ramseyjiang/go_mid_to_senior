package machinelearning

import "strings"

// Step 1: Tokenization
func tokenize(in <-chan string) <-chan []string {
	out := make(chan []string)
	go func() {
		for s := range in {
			out <- strings.Fields(s)
		}
		close(out)
	}()
	return out
}

// Step 2: Stop-word Removal
func removeStopWords(in <-chan []string) <-chan []string {
	stopWords := map[string]struct{}{
		"and": {},
		"the": {},
		"is":  {},
		"of":  {},
	}
	out := make(chan []string)
	go func() {
		for tokens := range in {
			filtered := make([]string, 0)
			for _, token := range tokens {
				if _, found := stopWords[token]; !found {
					filtered = append(filtered, token)
				}
			}
			out <- filtered
		}
		close(out)
	}()
	return out
}

// Step 3: Stemming
func stem(in <-chan []string) <-chan []string {
	// A very naive example just for illustration
	stemmingRules := map[string]string{
		"running": "run",
		"flies":   "fly",
	}

	out := make(chan []string)
	go func() {
		for tokens := range in {
			stemmed := make([]string, 0)
			for _, token := range tokens {
				if replacement, found := stemmingRules[token]; found {
					stemmed = append(stemmed, replacement)
				} else {
					stemmed = append(stemmed, token)
				}
			}
			out <- stemmed
		}
		close(out)
	}()
	return out
}

// Step 4: Vectorization (Placeholder)
func vectorised(in <-chan []string) <-chan []float64 {
	// Dummy vectorization
	out := make(chan []float64)
	go func() {
		for _ = range in {
			out <- []float64{0.1, 0.2, 0.3} // Dummy values
		}
		close(out)
	}()
	return out
}
