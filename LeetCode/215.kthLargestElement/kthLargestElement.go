// stats:
// Runtime: 76ms - Beats 16.78%
// Memory: 11.47MB - Beats 29.57%

package main

import "fmt"

func main() {
	fmt.Println(kthLargestElement([]int{3, 2, 1, 5, 6, 4}, 2))                       // 5
	fmt.Println(kthLargestElement([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))              // 4
	fmt.Println(kthLargestElement([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 10, 8, 9, 7}, 4)) // 7
}

func kthLargestElement(nums []int, k int) int {
	insertIntoMaxHeap := func(maxHeap []int, num int) []int {
		maxHeap = append(maxHeap, num)
		i := len(maxHeap) - 1
		for i > 0 {
			parentIndex := (i - 1) / 2

			if maxHeap[i] > maxHeap[parentIndex] {
				maxHeap[i], maxHeap[parentIndex] = maxHeap[parentIndex], maxHeap[i]
				i = parentIndex
			} else {
				break
			}
		}

		return maxHeap
	}

	popFromMaxHeap := func(maxHeap []int) []int {
		iLast := len(maxHeap) - 1
		maxHeap[iLast], maxHeap[0] = maxHeap[0], maxHeap[iLast]
		maxHeap = maxHeap[:iLast]
		i := 0
		for {
			iLeftChild := 2*i + 1
			iRightChild := 2*i + 2
			currNode := i

			if iLeftChild < len(maxHeap) && maxHeap[iLeftChild] > maxHeap[currNode] {
				currNode = iLeftChild
			}

			if iRightChild < len(maxHeap) && maxHeap[iRightChild] > maxHeap[currNode] {
				currNode = iRightChild
			}

			if currNode == i {
				break
			}

			maxHeap[i], maxHeap[currNode] = maxHeap[currNode], maxHeap[i]
			i = currNode
		}

		return maxHeap
	}

	if len(nums) == 0 {
		return 0
	}

	// build max heap and ignore duplicate nums
	maxHeap := []int{}
	for _, num := range nums {
		maxHeap = insertIntoMaxHeap(maxHeap, num)
	}

	// pop k-1 elems
	for _ = range k - 1 {
		maxHeap = popFromMaxHeap(maxHeap)
	}
	return maxHeap[0]
}
