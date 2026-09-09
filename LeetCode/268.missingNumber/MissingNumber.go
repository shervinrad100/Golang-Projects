// stats:
// Runtime:0ms, Beats 100.00%
// Memory: 20.42MB, Beats 41.55%
package main

import "fmt"

func main() {
	fmt.Println(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func MissingNumber(nums []int) int {
	numsLength := len(nums)
	runningSum := 0
	for _, n := range nums {
		runningSum += n
	}
	return int(numsLength*(numsLength+1)/2 - runningSum)
}
