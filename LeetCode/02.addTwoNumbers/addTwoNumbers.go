// stats:
// Runtime: 4 ms, faster than 97.65%
// Memory: 	4.7 MB, less than 47.01%

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceFromListNode(L *ListNode) []int {
	ans := []int{}
	for {
		ans = append(ans, L.Val)
		if L.Next != nil {
			L = L.Next
		} else {
			break
		}
	}
	return ans
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Pointer := l1
	l2Pointer := l2
	ansHead := ListNode{}
	ansPointer := &ansHead
	carry := 0
	for {
		sum := l1Pointer.Val + l2Pointer.Val + carry
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		ansPointer.Val = sum
		if (l1Pointer.Next == nil) && (l2Pointer.Next == nil) && (carry == 0) {
			break
		} else if (l1Pointer.Next == nil) && (l2Pointer.Next == nil) && (carry != 0) {
			ansPointer.Next = &ListNode{Val: carry}
			carry = 0
			break
		} else if (l1Pointer.Next != nil) && (l2Pointer.Next == nil) {
			l1Pointer = l1Pointer.Next
			l2Pointer = &ListNode{}
			ansPointer.Next = &ListNode{}
			ansPointer = ansPointer.Next
		} else if (l1Pointer.Next == nil) && (l2Pointer.Next != nil) {
			l2Pointer = l2Pointer.Next
			l1Pointer = &ListNode{}
			ansPointer.Next = &ListNode{}
			ansPointer = ansPointer.Next
		} else {
			l1Pointer = l1Pointer.Next
			l2Pointer = l2Pointer.Next
			ansPointer.Next = &ListNode{}
			ansPointer = ansPointer.Next
		}

	}

	return &ansHead
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 3},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{Val: 6,
			Next: &ListNode{Val: 4},
		},
	}

	output := AddTwoNumbers(&l1, &l2)
	fmt.Println(SliceFromListNode(output))

}
