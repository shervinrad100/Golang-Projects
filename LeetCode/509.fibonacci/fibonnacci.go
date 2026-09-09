package main

import (
	"fmt"
	"time"
)

func main() {
	// recursive no memoization
	fmt.Println("********* recursive no memoization *********")
	n := 5
	start := time.Now()
	fib := fibonnacciRecursive(n)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))

	n = 50 // grinds to a halt
	start = time.Now()
	fib = fibonnacciRecursive(n)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))

	// recursive no memoization
	fmt.Println("********* recursive with memoization *********")
	cache := map[int]int{}
	n = 5
	start = time.Now()
	fib = fibonnacciRecursiveMemo(n, &cache)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))

	n = 50
	start = time.Now()
	fib = fibonnacciRecursiveMemo(n, &cache)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))

	// iterative
	fmt.Println("********* iterative *********")
	n = 5
	start = time.Now()
	fib = fibonnacci(n)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))

	n = 50
	start = time.Now()
	fib = fibonnacci(n)
	fmt.Printf("fib(%d) = %d, calculated in %v\n", n, fib, time.Since(start))
}

func fibonnacciRecursive(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonnacciRecursive(n-1) + fibonnacciRecursive(n-2)
}

func fibonnacciRecursiveMemo(n int, fibList *map[int]int) int {

	if fib, ok := (*fibList)[n]; ok {
		return fib
	}
	if n <= 2 {
		(*fibList)[n] = 1
	} else {
		(*fibList)[n] = fibonnacciRecursiveMemo(n-1, fibList) + fibonnacciRecursiveMemo(n-2, fibList)
	}
	return (*fibList)[n]
}

func fibonnacci(n int) int {
	ans := 0
	fibList := map[int]int{}
	for i := range n + 1 {
		if i <= 2 {
			fibList[i] = 1
			ans += 1
		} else {
			fibList[i] = fibList[i-1] + fibList[i-2]
		}
	}
	return fibList[n]
}
