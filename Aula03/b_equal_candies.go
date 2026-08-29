package main

import (
	"fmt"
	"slices"
)

func abs(v int) int{
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	var cases int
	fmt.Scan(&cases)

	for range cases {
		var boxSize int
		fmt.Scan(&boxSize)
		candies := make([]int, boxSize)
		for i := range boxSize {
			var candie int
			fmt.Scan(&candie)
			candies[i] = candie
		}
		var sum int
		delta := slices.Min(candies)
		for i := range candies{
			if i+1 < len(candies) && candies[i] - candies[i+1] != 0{
				if candies[i] > delta {
					sum += candies[i]-delta
				} 
			} else if candies[i] != delta {
				sum += candies[i]-delta
			}
		}

		fmt.Printf("%d\n", sum)
	}
}
