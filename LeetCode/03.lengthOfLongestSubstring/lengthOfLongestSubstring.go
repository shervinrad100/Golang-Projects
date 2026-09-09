// stats:
// Runtime: 4 ms, faster than 90.80%
// Memory: 	3.1 MB, less than 58.27%

package main

import "fmt"

func LengthOfLongestSubstring(s string) int {

	n := len(s)
	ans := 0
	var m map[byte]int = map[byte]int{}

	for j, i := 0, 0; j < n; j++ {
		if val, found := m[s[j]]; found {
			if val > i {
				i = val
			}
		}

		if ans < j-i+1 {
			ans = j - i + 1
		}
		m[s[j]] = j + 1

	}

	return ans
}

func main() {
	var testCase string = "abcabcbb"
	output := LengthOfLongestSubstring(testCase)
	fmt.Println(output)

}
