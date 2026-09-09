// stats:
// Runtime: 0ms - Beats 100.00%
// Memory: 3.96 MB - Beats 54.34%

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
}

func wordPattern(pattern, s string) bool {
	stringList := strings.Split(s, " ")
	if len(stringList) != len(pattern) {
		return false
	}

	wordMap := map[string]byte{} // map[word]token
	charMap := map[byte]string{} // map[token]string

	for i, word := range stringList {

		if token, ok := wordMap[word]; ok {
			if pattern[i] != token {
				return false
			}
		} else {
			wordMap[word] = pattern[i]
		}

		if w, ok := charMap[pattern[i]]; ok {
			if w != word {
				return false
			}
		} else {
			charMap[pattern[i]] = word
		}
	}

	return true
}
