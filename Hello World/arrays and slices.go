package main

import (
	"fmt"
)

var (
	identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
)

func main() {
	// array of size 3 holding ints and you fill it with values 1 to 3
	// grades = [3]int
	// grades[0] = 1
	// grades[1] = 2
	// grades[2] = 3
	grades := [...]int{1, 2, 3}
	fmt.Println("Grades:", grades)
	fmt.Printf("first grade: %v (type %T)\n", grades[0], grades[0])

	fmt.Println(identityMatrix)

	// arrays in Go are not referencing where the internal values are
	// its actually stored as a value
	// so if you have an array with len(array) << 10 then it copies the whole thing across and is v inefficient
	// you can create a reference by having b:= &a which makes b point to a

	// slices dont have the fixed length that arrays do
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	// slice is the projection of an array, capacity shows the len of the underlying array cap(a)
	// slices are reference types unlike arrays
	// you can use the make function to create a slice with larger capacity

	a := make([]int, 3, 100) // len 3 cap 100
	fmt.Println(a)
	fmt.Printf("len %v, cap %v\n", len(a), cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("len %v, cap %v\n", len(a), cap(a))

	// when you append to a slice and with that exceed the capacity,
	// the slice will be copied over to an array and this can be very expensive for large arrays/slices

	b := []int{}
	fmt.Println(b)
	fmt.Printf("len %v, cap %v\n", len(b), cap(b))

	b = append(b, 1)
	fmt.Println(b)
	fmt.Printf("len %v, cap %v\n", len(b), cap(b))

	b = append(b, 2, 3, 4)
	// b = append(b, []int{5,6,7}) will get error
	b = append(b, []int{5, 6, 7}...) // ... is like the * operator in python and will spread the iterable

	// removing elements
	fmt.Println(b)
	b = append(b[:4], b[5:]...)
	// c := append(b[:4], b[5:]...)
	//fmt.Println(c) // the underlying array changes if you assign it to another slice
	fmt.Println(b)
}
