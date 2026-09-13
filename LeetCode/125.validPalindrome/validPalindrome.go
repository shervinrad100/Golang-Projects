package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isValidPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isValidPalindrome("race a car"))                     // false
	fmt.Println(isValidPalindrome(" "))                              // true
}

func isValidPalindrome(s string) bool {
	// shrinking window
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		// if it's not alphanumeric pass
		for !unicode.IsLetter(rune(s[p1])) && !unicode.IsNumber(rune(s[p1])) {
			p1++
		}
		for !unicode.IsLetter(rune(s[p2])) && !unicode.IsNumber(rune(s[p2])) {
			p2--
		}

		if unicode.ToLower(rune(s[p1])) != unicode.ToLower(rune(s[p2])) {
			return false
		}
		p1++
		p2--
	}
	return true
}
