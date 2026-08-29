package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	original := "hello"
	index := 0

	for _, c := range input {
		if index < len(original) && byte(c) == original[index] {
			index++
		}
	}

	if index == len(original) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
