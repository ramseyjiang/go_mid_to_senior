package ratelimit

import (
	"testing"
)

func TestIsAllow(t *testing.T) {
	clientID := 123
	clientID2 := 4
	for i := 0; i < 5; i++ {
		if IsAllow(clientID) {
			t.Log("In 5 times requests, the return is true.")
		} else {
			t.Error("Over 5 times requests, the return is false.")
		}
	}

	for i := 0; i < 5; i++ {
		if IsAllow(clientID2) {
			t.Log("In 5 times requests, the return is true.")
		} else {
			t.Error("Over 5 times requests, the return is false.")
		}
	}
}
