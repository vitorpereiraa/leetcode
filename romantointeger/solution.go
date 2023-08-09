package main

import "fmt"

func main() {
	var romanNumeral string

	fmt.Print("Insert a roman numeral: ")
	fmt.Scan(&romanNumeral)

	var converted int = romanToInt(romanNumeral)

	fmt.Printf("Converted to integer: %d", converted)
}

func romanToInt(s string) int {
	var romanNumerals = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var converted int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanNumerals[s[i+1]] > romanNumerals[s[i]] {
			converted -= romanNumerals[s[i]]
		} else {
			converted += romanNumerals[s[i]]
		}
	}

	return converted
}
