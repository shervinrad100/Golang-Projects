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

	var backtrack func(i int, decision string)
	backtrack = func(i int, decision string) {
		if i == len(nums) {

			// WRONG!!
			// because sub lives in the heap. append(ans, sub) will add the pointer to the heap.
			// You can see that the address to sub stays the same (it changes when the 2kb memory needs to be reallocated and grown)
			// ans = append(ans, sub)
			// fmt.Println("i", i, "decision", decision)
			// fmt.Printf("ans: %v at %p\nsub: %v at %p\n", ans, ans, sub, sub)

			// CORRECT!!
			temp := make([]int, len(sub))
			copy(temp, sub)
			ans = append(ans, temp)
			return
		}

		// decision 1: dont add
		backtrack(i+1, "dont add")

		// decision 2: add num to sub
		sub = append(sub, nums[i])
		backtrack(i+1, "add")
		sub = sub[:len(sub)-1]
	}
	backtrack(0, "")

	// for i, slice := range ans {
	// 	fmt.Printf("ans[%d] value: %v | Underlying array address: %p\n", i, slice, slice)
	// }
	/*
		Look at the debug. The items in ans are pointing at the same memory address
			ans[0] value: [] | Underlying array address: 0x1031503a0
			ans[1] value: [2] | Underlying array address: 0x140000a8028
			ans[2] value: [2] | Underlying array address: 0x140000a8028
			ans[3] value: [1 2] | Underlying array address: 0x140000a8070
			ans[4] value: [1] | Underlying array address: 0x140000a8070
			ans[5] value: [1 2] | Underlying array address: 0x140000a8070
			ans[6] value: [1 2] | Underlying array address: 0x140000a8070
			ans[7] value: [1 2 3] | Underlying array address: 0x140000d0020
	*/

	return ans
}
