package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	// we want to print the error log in case some error happens
	defer func() { // this is called an anonymous function. its a throw away function
		if err := recover(); err != nil { // initialise an err variable by calling the recover function
			log.Println("Error:", err)
		}
	}() // you call the anonymous function
	panic("something bad happened")
	fmt.Println("done panicking")
}

// so we can see that panicking allows functions higher in the call stack to finish
// and will only terminate the current function
