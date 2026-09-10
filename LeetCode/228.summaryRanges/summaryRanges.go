// stats:
// Runtime: 0ms - Beats 100.00%
// Memory: 3.80MB - Beats 99.63%

package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0","2->4","6","8->9"]
	fmt.Println(summaryRanges([]int{}))                    // []
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		// this if loop saves us 0.3MB
		return []string{}
	}
	output := []string{}
	p1, p2 := 0, 0

	for p2 < len(nums) {
		if p2+1 == len(nums) || nums[p2+1]-nums[p2] != 1 {
			if p1 == p2 {
				output = append(output, fmt.Sprintf("%d", nums[p1]))
			} else {
				output = append(output, fmt.Sprintf("%d->%d", nums[p1], nums[p2]))
			}
			p1 = p2 + 1
		}
		p2++
	}
	return output
}
