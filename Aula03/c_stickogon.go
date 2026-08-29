package main

import (
	"fmt"
)

func main() {
	var cases int
	fmt.Scan(&cases)
	sticksMap := make(map[int]int)
	
	for range cases{
		var isPolygon int 
		var sticks int
		fmt.Scan(&sticks)
		for range sticks {
			var stick int
			fmt.Scan(&stick)
			sticksMap[stick]++
			if sticksMap[stick] >= 3{
				isPolygon++
				delete(sticksMap, stick)
			}
		}
		clear(sticksMap)
		fmt.Printf("%d\n", isPolygon)
	}

}

