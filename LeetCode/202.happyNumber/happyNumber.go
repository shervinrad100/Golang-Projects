package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(happyNumber(19)) // true
	fmt.Println(happyNumber(2))  // false
}

func happyNumber(n int) bool {
	sumOfDigitsSquared := func(digits [10]int) int {
		runningSum := 0
		for i, multiple := range digits {
			runningSum += i * i * multiple
		}
		return runningSum
	}

	seenCombinations := map[[10]int]bool{}

	for {
		// turn the number into usable array
		digits := strings.Split(fmt.Sprintf("%d", n), "")

		// create the seen set
		tmpMap := [10]int{}
		for _, digit := range digits {
			var d int
			fmt.Sscanf(digit, "%d", &d)
			tmpMap[d] += 1
		}

		// make a decision on the combination
		n = sumOfDigitsSquared(tmpMap)
		if seenCombinations[tmpMap] {
			if n != 1 {
				return false
			} else {
				return true
			}
		} else {
			seenCombinations[tmpMap] = true
		}
	}

}
