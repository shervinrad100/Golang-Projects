package main

import "fmt"

func main() {
	fmt.Println("************** Recursive with memo **************")
	memo := make(map[int]int)
	fmt.Println(fewestCoinsRecursive(13, []int{1, 4, 5}, &memo))

}

type Number interface {
	~int | ~float32 | ~float64
}

func min[N Number](nums ...N) (N, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	_min := nums[0]
	for _, num := range nums[1:] {
		if num < _min {
			_min = num
		}
	}
	return _min, true
}

func minWithNils(a, b *int) int {
	if a == nil {
		return *b
	}
	if b == nil {
		return *a
	}
	_min, _ := min(*a, *b)
	return _min

}

func fewestCoinsRecursive(target int, coinSet []int, cache *map[int]int) (int, bool) {
	if ans, ok := (*cache)[target]; ok {
		return ans, true
	}
	if target == 0 {
		return 0, true
	}
	ans := 0
	hasAnswer := false
	for _, coin := range coinSet {
		subProblem := target - coin
		if subProblem < 0 {
			continue
		}
		subAnswer, valid := fewestCoinsRecursive(subProblem, coinSet, cache)
		if !valid {
			continue
		}
		currentCoins := subAnswer + 1

		if !hasAnswer || currentCoins < ans {
			ans = currentCoins
			hasAnswer = true
		}
	}
	(*cache)[target] = ans
	return ans, true
}
