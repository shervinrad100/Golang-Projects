// tests must be saved in a file named like <something>_test.go

package main

import (
	"fmt"
	"testing"
)

type Test struct {
	i      int
	input  []int
	target int
	ans    [2][2]int
}

func TestTwoSums(t *testing.T) { // test func must start with "Test" followed by a phrase that starts with capital letter

	tests := []Test{
		{
			i:      1,
			input:  []int{7, 2, 11, 15},
			target: 9,
			ans:    [2][2]int{[2]int{1, 0}, [2]int{0, 1}},
		},
		{
			i:      2,
			input:  []int{1, 3, 4, 2},
			target: 6,
			ans:    [2][2]int{[2]int{2, 3}, [2]int{3, 2}},
		},
	}

	for _, test := range tests {
		output := TwoSums(test.input, test.target)
		var correct bool = false
		for _, indexList := range test.ans {
			if indexList == output {
				correct = true
			}
		}
		if !correct {
			t.Errorf("Test %d failed!\nExpected %d, Got %d", test.i, test.ans, output)
		} else {
			fmt.Printf("Testcase %v passed!\n", test.i)
		}
	}

}

// go test [--verbose|-v] <path_to_PACKAGE>
