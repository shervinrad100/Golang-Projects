package main

import (
	"fmt"
	"testing"
)

func listNodeFromSlice(slice *[]int) *ListNode {
	ans := ListNode{}
	ans.Val = (*slice)[0]
	NextItems := len(*slice)
	if NextItems != 1 {
		sliceOf := (*slice)[1:]
		ans.Next = listNodeFromSlice(&sliceOf)
	}
	return &ans
}

func sliceFromListNode(L *ListNode) []int {
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

func TestAddTwoNumbers(t *testing.T) {
	var testQuestions [][2][]int = [][2][]int{
		// test1
		[2][]int{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
		},
		// test2
		[2][]int{
			[]int{0},
			[]int{0},
		},

		// test3
		[2][]int{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
		},
	}

	var testAnswers [][]int = [][]int{
		// test1
		[]int{7, 0, 8},
		// test2
		[]int{0},
		// test3
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	}

	for testIndex, test := range testQuestions {
		l1 := listNodeFromSlice(&(test[0]))
		l2 := listNodeFromSlice(&(test[1]))
		output := sliceFromListNode(AddTwoNumbers(l1, l2))

		var correct bool = true

		for key, val := range output {
			if val != testAnswers[testIndex][key] {
				correct = false
				break
			}
		}

		if !correct {
			t.Errorf("Test %d failed!\nExpected output: %d\nYour output: %d",
				testIndex, testAnswers[testIndex], output,
			)
		} else {
			fmt.Printf("Testcase %v passed!\n", testIndex)
		}

	}

}
