// stats:
// Runtime: 29ms - Beats 95.22%
// Memory: 13.06MB - Beats 70.19%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}, 3))          // true
	fmt.Println(containsDuplicate([]int{1, 0, 1, 1}, 1))          // true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))    // false
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1, 2, 1, 3}, 1)) // false
}

func containsDuplicate(nums []int, k int) bool {
	mem := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := mem[num]; ok {
			if i-j <= k {
				return true
			}
		}
		mem[num] = i
	}
	return false
}
