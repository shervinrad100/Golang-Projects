package main

import "fmt"

func main() {
	fmt.Println(targetSum([]int{1, 1, 1, 1, 1}, 3))
}

func targetSum(nums []int, target int) int {
	return 0 // there's a more efficient way of doing this by computing the decision tree and getting only the desired path
}

func targetSumRecursive(nums []int, target int) int {
	var backtrack func(i, currSum int) int
	backtrack = func(i, currSum int) int {
		if i == len(nums) {
			if currSum == target {
				return 1
			} else {
				return 0
			}
		}
		return backtrack(i+1, currSum+nums[i]) + backtrack(i+1, currSum-nums[i])
	}
	return backtrack(0, 0)
}
