package main

import "fmt"

func main(){
	var input string
	fmt.Scan(&input)
	leftCount := 0
	rightCount := 0
	for i := range input {
		if string(input[i]) == "(" { rightCount++ }
		if string(input[i]) == ")" { leftCount++ }
	}
	if rightCount == leftCount {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}

