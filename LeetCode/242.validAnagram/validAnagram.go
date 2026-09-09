// stats:
// Runtime: 0ms - Beats 100.00%
//Memory: 4.74MB - Beats 58.50%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(validAnagram("anagram", "nagaram")) // true
	fmt.Println(validAnagram("rat", "car"))         // false
}

func validAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	englishLettersList := [26]int{}

	for i := 0; i < len(s); i++ {
		englishLettersList[s[i]-'a']++
		englishLettersList[t[i]-'a']--
	}
	for i := range englishLettersList {
		if englishLettersList[i] != 0 {
			return false
		}
	}
	return true

}
