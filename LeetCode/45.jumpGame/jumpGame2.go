package main

import (
	"fmt"
)

func main() {
	fmt.Println(jumpGame2([]int{1}))                                           // 0
	fmt.Println(jumpGame2([]int{0}))                                           // 0
	fmt.Println(jumpGame2([]int{2, 3, 1, 1, 4}))                               // 2
	fmt.Println(jumpGame2([]int{2, 3, 0, 1, 4}))                               // 2
	fmt.Println(jumpGame2([]int{2, 0, 2, 0, 1}))                               // 2
	fmt.Println(jumpGame2([]int{4, 1, 1, 3, 1, 1, 1}))                         // 2
	fmt.Println(jumpGame2([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3})) // 2
}

func jumpGame2(nums []int) int {
	jumps := 0
	p1, p2 := 0, 0 // sliding window
	var maxReach int

	for p2 < len(nums)-1 {
		maxReach = 0
		for i := p1; i <= p2; i++ {
			if maxReach < nums[i]+i {
				maxReach = nums[i] + i
			}
		}
		p1 = p2 + 1
		p2 = maxReach
		jumps++
	}

	return jumps
}
