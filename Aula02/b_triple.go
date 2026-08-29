package main

import (
	"fmt"
)


func main() {
	var tests int
	fmt.Scan(&tests)
	for range tests {
		var inputLength int
		var hasRepeated bool
		var valueRepeted int

		fmt.Scan(&inputLength)
		count := make([]int, inputLength+1)
		
		for range inputLength {
			var input int
			fmt.Scan(&input)
			
			count[input]++

			if count[input] >= 3 {
				hasRepeated = true
				valueRepeted = input
			}
		}
		if !hasRepeated {
			fmt.Printf("-1\n")
		} else {
			fmt.Printf("%d\n", valueRepeted)
		}
	}
}
