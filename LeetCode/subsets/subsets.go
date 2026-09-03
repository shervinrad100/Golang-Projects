package main

import "fmt"

func main() {
	fmt.Println(subsetsRecursive([]int{1, 2, 3}))
	fmt.Println(subsetsRecursive([]int{0}))
	// fmt.Println(subsets([]int{1, 2, 3}))
	// fmt.Println(subsets([]int{0}))
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

func subsetsRecursive(nums []int) any {
	ans, sub := [][]int{}, []int{}

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			// ans = append(ans, sub[:]) WRONG!! Why?
			temp := make([]int, len(sub))
			copy(temp, sub)
			ans = append(ans, temp)
			return
		}

		// decision 1: dont pick
		backtrack(i + 1)

		// decision 2: pick num
		sub = append(sub, nums[i])
		backtrack(i + 1)
		sub = sub[:len(sub)-1]
	}
	backtrack(0)
	return ans
}
