// stats :
// Runtime: 3ms - Beats 93.37%
// Memory: 6.29MB - Beats 39.38%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasCycle(makeList([]int{3, 2, 0, -4}, 1))) // true
	fmt.Println(hasCycle(makeList([]int{1, 2}, 0)))        // true
	fmt.Println(hasCycle(makeList([]int{1}, -1)))          // false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(nodes []int, pos int) *ListNode {
	nodeRef := []*ListNode{}

	root := &ListNode{
		Val: nodes[0],
	}
	nodeRef = append(nodeRef, root)

	for i := 1; i < len(nodes); i++ {
		node := &ListNode{
			Val: nodes[i],
		}
		nodeRef[i-1].Next = node
		nodeRef = append(nodeRef, node)
	}

	if pos > 0 {
		nodeRef[len(nodeRef)-1].Next = nodeRef[pos]
	}

	return root
}

func hasCycle(head *ListNode) bool {
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if tortoise == hare {
			return true
		}
	}

	return false

}
