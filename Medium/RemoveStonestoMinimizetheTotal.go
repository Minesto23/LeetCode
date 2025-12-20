package main

import (
	"fmt"
)

// minStoneSum removes half of the largest pile k times
// and returns the sum of the remaining stones.
func minStoneSum(piles []int, k int) int {
	for j := 0; j < k; j++ {
		// Find index of the largest pile
		maxIdx := 0
		for i := 1; i < len(piles); i++ {
			if piles[i] > piles[maxIdx] {
				maxIdx = i
			}
		}

		// Remove half of the largest pile
		piles[maxIdx] -= piles[maxIdx] / 2
	}

	// Sum remaining piles
	sum := 0
	for _, value := range piles {
		sum += value
	}

	return sum
}

func main() {
	piles := []int{5, 4, 9}
	k := 2

	result := minStoneSum(piles, k)
	fmt.Println("Minimum sum after", k, "operations:", result)
}
