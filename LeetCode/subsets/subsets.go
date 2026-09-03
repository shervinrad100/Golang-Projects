package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}

func subsets(nums []int) []string {
	ans := map[string]bool{}
	for p1 := range len(nums) {
		p2 := p1
		for p2 <= len(nums) {
			ans[fmt.Sprint(nums[p1:p2])] = true
			p2++
		}
	}
	keys := []string{}
	for k := range ans {
		keys = append(keys, k)
	}
	return keys
}
