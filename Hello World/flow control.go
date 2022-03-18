package main

import (
	"fmt"
)

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func main() {
	// you can have an initialiser in the if statement like so ...
	hashMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	if val, inMap := hashMap[1]; inMap {
		fmt.Println(val)
	}
	// || or
	// && and
	// ! not

	// short circuiting:
	// if you have a few logical tests, the first one that returns true will stop the rest of the
	// tests form running
	input := 5

	if input == 5 || returnTrue() {
		fmt.Println("if block ran")
		// you wont see "returning true"
	}

	if returnTrue() {
		fmt.Println("if block ran")
	}

	if false && returnTrue() {
		fmt.Println("Youll see none of this")
	}

	// switch statements compares the value that you give to many other values
	// so you dont have to rewrite the logic comparison multiple times in the if block
	switch num := 100; num { // you can innitialise the tag or pass a variable
	case 10:
		fmt.Println("input is equal to 10")
	case 1, 2, 3:
		fmt.Println("input is equal to 1,2 or 3")
	default:
		fmt.Println("none of the above")
	}
	// your tags cannot overlap

	// you can also have a switch statement with no tag:
	switch {
	case input <= 10:
		fmt.Println("the input is leq to 10")
	case input <= 20:
		fmt.Println("the input is leq to 10")
	default:
		fmt.Println("none of the above")
	}
	// your switch statemet will break after the first case is met unless you use the keyword "fallthrough"
	switch {
	case input <= 10:
		fmt.Println("the input is leq to 10")
		fallthrough // this will cause the next block to run regardless of the logic
	case input > 100: // input <= 20:
		fmt.Println("the input is leq to 10")
	default:
		fmt.Println("none of the above")
	}

	// you can assign anything to an interface var type
	var i interface{} = [2]int{}
	switch i.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	case [3]int:
		fmt.Println("is [3]int")
	default:
		fmt.Println("is other")
		break
		fmt.Print("this wont print")
	}

	// for loops

	for i := 0; i < 5; i++ { // i+= 2 will increment i by 2
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 2; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	k := 0
	for ; k < 2; k++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}

	for k < 5 {
		fmt.Println(i)
		k++
	}

	for {
		fmt.Println(k)
		k++
		if k > 5 {
			break
		}
	}
	// you also have "continue" in go

	// if youre in a nested loop and want to break out of the outer
	// loop you can label your loop and break out of it by refering to that loop

	fmt.Println("loop tag:...")
Loop1:
	for k := 0; true; {
		for j := 0; j < 5; j++ {
			fmt.Println(k, j)
			if k > 2 {
				break Loop1
			}
		}
		k++
	}

	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println("index", index, "value", value)
	}

}
