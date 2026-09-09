// 986 / 987 test cases passed.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func readTestFile() []byte {
	content, err := ioutil.ReadFile("edgeTestCase.txt")
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func TestLengthOfLongestSubstring(t *testing.T) {

	type testCase struct {
		testString     string
		expectedOutput int
	}

	testCases := []testCase{
		{
			testString:     "",
			expectedOutput: 0,
		},
		{
			testString:     " ",
			expectedOutput: 1,
		},
		{
			testString:     "abcabcbb",
			expectedOutput: 3,
		},
		{
			testString:     "bbbbb",
			expectedOutput: 1,
		},
		{
			testString:     "pwwkew",
			expectedOutput: 3,
		},
		{
			testString:     string(readTestFile()),
			expectedOutput: 95,
		},
	}

	for testIndex, test := range testCases {
		output := LengthOfLongestSubstring(test.testString)
		if output != test.expectedOutput {
			t.Errorf("Test %d failed!\nExpected output: %d\nYour output: %d\n",
				testIndex+1, test.expectedOutput, output)
		} else {
			fmt.Printf("Testcase %v passed!\n", testIndex+1)
		}
	}

}
