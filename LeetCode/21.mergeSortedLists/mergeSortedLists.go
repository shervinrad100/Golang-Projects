// stats
// Runtime: 0ms - Beats 100.00%
// Memory: 4.46MB - Beats 30.70%

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeSortedLists(makeLinkedList([]int{1, 2, 4}), makeLinkedList([]int{1, 3, 4}))) // [1,1,2,3,4,4]
	fmt.Println(mergeSortedLists(makeLinkedList([]int{}), makeLinkedList([]int{})))               // []
	fmt.Println(mergeSortedLists(makeLinkedList([]int{}), makeLinkedList([]int{0})))              // [0]
	fmt.Println(mergeSortedLists(makeLinkedList([]int{-9, 3}), makeLinkedList([]int{5, 7})))      // [-9,3,5,7]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (L *ListNode) String() string {
	if L == nil {
		return "[]"
	}
	current := L
	var s strings.Builder
	fmt.Fprintf(&s, "[%d", current.Val)
	for current.Next != nil {
		current = current.Next
		fmt.Fprintf(&s, ",%d", current.Val)
	}
	s.WriteString("]")
	return s.String()
}

func makeLinkedList(lst []int) *ListNode {
	if len(lst) == 0 {
		return nil
	}
	root := &ListNode{
		Val: lst[0],
	}
	current := root
	for i := 1; i < len(lst); i++ {
		node := &ListNode{
			Val: lst[i],
		}
		current.Next = node
		current = node
	}
	return root
}

func mergeSortedLists(l1, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	mergedListRoot := &ListNode{}
	current := mergedListRoot

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			current.Next = &ListNode{
				Val: p1.Val,
			}
			p1 = p1.Next
		} else {
			current.Next = &ListNode{
				Val: p2.Val,
			}
			p2 = p2.Next
		}
		current = current.Next
	}

	if p1 == nil {
		// append p2 to ans
		for p2 != nil {
			current.Next = &ListNode{
				Val: p2.Val,
			}
			p2 = p2.Next
			current = current.Next
		}
	}
	if p2 == nil {
		// append p1 to ans
		for p1 != nil {
			current.Next = &ListNode{
				Val: p1.Val,
			}
			p1 = p1.Next
			current = current.Next
		}
	}

	return mergedListRoot.Next
}
