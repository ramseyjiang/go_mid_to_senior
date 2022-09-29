package twosum

import (
	"log"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	if res := twoSum1(nums, target); len(res) > 0 {
		log.Println("target is the sum of two index positions", res)
	} else {
		t.Error("Target is not any two num sum in the nums.")
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	if res := twoSum2(nums, target); len(res) > 0 {
		log.Println("target is the sum of two index positions", res)
	} else {
		t.Error("Target is not any two num sum in the nums.")
	}
}
