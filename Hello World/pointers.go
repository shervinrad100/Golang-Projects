package main

import (
	"fmt"
)

// here we create our own custom variable type in the form of a struct
type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a // b is now pointing to a instead of copying it across
	// *int creates a pointer to an integer
	fmt.Println(a, b) // b prints the memory location of a
	// the & operator gives the address of an object so you can see that the two values are equal
	fmt.Println(&a, b)
	// the * is the opposite of that and dereferences the address and gives the value
	fmt.Println(a, *b)

	// https://www.youtube.com/watch?v=sTFJtxJXkaY
	// the variable m is going to hold an address to an object of type myStruct
	// to reiterate, var m is a pointer var
	// so to assign a value to m, it should be an address of a variable
	// we declare a variable of type myStruct with value {foo: 42} and take the address of that with the & operator
	var m *myStruct
	m = &myStruct{foo: 42}
	fmt.Println(m)

	// you will learn this later but this is used when you want to edit variables outside of the active frame
	// if you return a referenced type, after the active frame is serviced by garbage collector
	// the underlying variable that holds the data will be deleted and so your reference is broken
	// to tackle this the compiler will save the value in a heap so that the reference is not broken
	// if you allocate too many variables in the heap
	//it can impact performance as it becomes cumbersome for the garbage collector

}
