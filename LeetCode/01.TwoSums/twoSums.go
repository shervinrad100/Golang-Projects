// stats:
// Runtime: 7 ms, faster than 76.34%
// Memory: 	4.3 MB, less than 31.84%

package main

import "fmt"

func TwoSums(nums []int, target int) [2]int {

	var hashMap map[int]int = map[int]int{}

	// iterate through the list and add it to a hash map
	for i := 0; i < len(nums); i++ {
		// check to see if complement has been added to hash table
		index, inMap := hashMap[target-nums[i]]
		if inMap {
			return [2]int{i, index}
		} else { // add index to hash map
			hashMap[nums[i]] = i
		}
	}

	return [2]int{}

}

func main() {
	fmt.Println(TwoSums([]int{7, 2, 11, 15}, 9))
}
