package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumNonOverlappingSubArrays([]int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15}, 3, 5)) // 109
	fmt.Println(sumNonOverlappingSubArrays([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2))          // 20
	fmt.Println(sumNonOverlappingSubArrays([]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3))       // 31
	fmt.Println(sumNonOverlappingSubArrays([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2))          // 29
}

type Number interface {
	~int | ~float32 | ~float64
}

func sum[N Number](n []N) N {
	var runningSum N
	for _, x := range n {
		runningSum += x
	}
	return runningSum
}

type dict struct {
	key   int
	value int
}
type byValueReverse []dict

func (v byValueReverse) Len() int {
	return len(v)
}
func (v byValueReverse) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v byValueReverse) Less(i, j int) bool {
	return v[i].value > v[j].value
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sumNonOverlappingSubArrays(nums []int, len1, len2 int) int {
	max1, max2 := []dict{}, []dict{}
	lenNums := len(nums)

	// find the sums and put them in memory
	for i := 0; i <= lenNums; i++ {
		if i+len1 <= lenNums {
			max1 = append(max1, dict{i, sum(nums[i : i+len1])})
		}
		if i+len2 <= lenNums {
			max2 = append(max2, dict{i, sum(nums[i : i+len2])})
		}
	}

	// sort the sums by DESC order
	sort.Sort(byValueReverse(max1))
	sort.Sort(byValueReverse(max2))

	// go through the sums and take the largest ones as long as they're not overlapping
	ans := 0
	for i := range max1 {
		for j := range max2 {
			if max2[j].key+len2 <= max1[i].key || max2[j].key >= max1[i].key+len1 {
				ans = max(ans, max1[i].value+max2[j].value)
			} else {
				continue
			}
		}
	}
	return ans
}
