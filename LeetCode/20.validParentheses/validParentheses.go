// stats
// Runtime: 0ms - Beats 100.00%
// Memory: 4.10MB - Beats 85.18%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(validParentheses("()"))     // true
	fmt.Println(validParentheses("()[]{}")) // true
	fmt.Println(validParentheses("(]"))     // false
	fmt.Println(validParentheses("([])"))   // true
	fmt.Println(validParentheses("([)]"))   // false

}

func validParentheses(s string) bool {
	stack := []byte{}
	pMapping := map[byte]byte{'}': '{', ']': '[', ')': '('}
	for i := range s {
		if open, ok := pMapping[s[i]]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == open {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false

}
