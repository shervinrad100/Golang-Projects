// stats :
// Runtime 96ms - Beats 65.38%
// Memory 15.11MB - Beats 76.92%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(lonelyNumebr([]int{10, 6, 5, 8}))
	fmt.Println(lonelyNumebr([]int{1, 3, 5, 3}))
}

func lonelyNumebr(nums []int) []int {
	mapped := map[int]int{}
	for _, num := range nums {
		mapped[num] += 1
	}

	ans := []int{}
	for k, v := range mapped {
		if _, ok := mapped[k-1]; ok {
			continue
		}
		if _, ok := mapped[k+1]; ok {
			continue
		}
		if v > 1 {
			continue
		}
		ans = append(ans, k)
	}
	return ans
}
