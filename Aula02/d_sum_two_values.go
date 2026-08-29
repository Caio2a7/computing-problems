package main

import (
	"fmt"
)

func main(){
	var size int
	var targetSum int

	fmt.Scan(&size, &targetSum)
	sums := make(map[int]int)

	var twoSum [2]int
	var hasTwoSum bool
	for i := range size {
		var input int
		fmt.Scan(&input)
		delta := targetSum-input 
		if index, ok := sums[delta]; ok {
			twoSum[0] = i+1
			twoSum[1] = index+1
			hasTwoSum = true
		}
		sums[input] = i
	}
	if hasTwoSum == true {
		fmt.Printf("%d %d\n", twoSum[0], twoSum[1])
	} else {
		fmt.Printf("-1\n")
	}

}

