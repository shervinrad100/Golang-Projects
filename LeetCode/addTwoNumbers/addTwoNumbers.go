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
		}
		ansPointer.Val = sum
		if (l1Pointer.Next == nil) && (l2Pointer.Next == nil) {
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
		Val:  11,
		Next: &ListNode{Val: 12},
	}
	l2 := ListNode{
		Val:  21,
		Next: &ListNode{Val: 22},
	}

	output := AddTwoNumbers(&l1, &l2)
	fmt.Println(output)
	fmt.Println(SliceFromListNode(output))
}
