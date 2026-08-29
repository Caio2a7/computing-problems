package main

import "fmt"

func main(){
	var cases int
	fmt.Scan(&cases)
	for range cases{
		var input int
		fmt.Scan(&input)
		var division int
		switch {
			case input >= 1900:
				division = 1
			case input >= 1600:
				division = 2
			case input >= 1400:
				division = 3
			default:
				division = 4
		}
		fmt.Printf("Division %d\n", division)
	}


}
