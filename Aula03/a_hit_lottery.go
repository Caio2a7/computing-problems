package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	billsOrder := []int{100,20,10,5,1}
	billSum := 0
	var tmp int = input 
	for _, bill := range billsOrder {
		billSum += tmp / bill
		tmp = input % bill
	}
	fmt.Printf("%d\n", billSum)
}

