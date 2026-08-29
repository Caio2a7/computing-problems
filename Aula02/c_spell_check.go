package main

import (
	"fmt"
	"unicode"
	"strings"
)

func main(){
	var testCases int 
	var fullName string = "Timur"
	fmt.Scan(&testCases)
	
	var nameLength int
	var name string
	hasUpperT := make([]bool, testCases)
	hasUpper := make([]bool, testCases)
	isDifferent := make([]bool, testCases)
	isRepeted := make([]bool, testCases)
	results := make([]string, testCases)
	for i := range testCases {
		fmt.Scan(&nameLength, &name)
		for j := range nameLength {
			if !strings.Contains(fullName, string(name[j])) {
				results[i] = "NO"
			}
			if strings.Count(name, string(name[j])) > 1{
				isRepeted[i] = true
			}
			if nameLength != len(fullName) {
				results[i] = "NO"
			}
			if string(name[j]) == "T" {
				if hasUpperT[i] {
					hasUpper[i] = true
				} else {
					hasUpperT[i] = true
				}
			} else if unicode.IsLetter(rune(name[j])) && unicode.IsUpper(rune(name[j])) {
				if strings.Contains(fullName, string(name[j])) != true{
					isDifferent[i] = true
				}
				hasUpper[i] = true
			}
		}
		if results[i] != "NO" && hasUpperT[i] == true && hasUpper[i] == false && isDifferent[i] == false && isRepeted[i] == false{
			results[i] = "YES"
		} else{
			results[i] = "NO"
		}
	}
	for _, result := range results{
		fmt.Printf("%s\n", result)
	}
}
