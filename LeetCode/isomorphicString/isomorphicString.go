// stats:
// Runtime: 1ms - Beats 61.55%
// Memory: 4.46MB - Beats 87.97%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isomorphicString("egg", "add"))     // true
	fmt.Println(isomorphicString("f11", "b23"))     // false
	fmt.Println(isomorphicString("paper", "title")) // true
	fmt.Println(isomorphicString("egcd", "adfd"))   // false
}

func isomorphicString(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_tMap := map[byte]byte{} // map[sByte]tByte
	t_sMap := map[byte]byte{} // map[tByte]sByte
	for i := range s {

		if tByte, ok := s_tMap[s[i]]; ok {
			if tByte != t[i] {
				return false
			}
		} else {
			s_tMap[s[i]] = t[i]
		}

		if sByte, ok := t_sMap[t[i]]; ok {
			if sByte != s[i] {
				return false
			}
		} else {
			t_sMap[t[i]] = s[i]
		}
	}
	return true
}
