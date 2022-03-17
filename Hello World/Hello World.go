package main

import (
	"fmt"
	"strconv"
)

// you can define variables for the whole package here
// also instead of repeating var a milliion times you can have them in a var block like so:
//var (
//	actorName    string = "Liam Neeson"
//	companion    string = "James Bond"
//	doctorNumber int    = 3
//	season       int    = 5
//)
//
// you can organise your code and have separate var blocks to group different variable categories
//var (
//	another   string
//	set       string
//	of        string
//	variables string
//)

// scopes of vars:
// first letter uppercase variables can be used globally while lower case ones are only for packages
// variables declared in a function cannot be called outside of that function

// use uppercase for all acronyms

func main() {
	fmt.Println("Hello World!")
	var i int // you can declare the type of the var without assignment
	i = 42
	fmt.Println("i =", i)
	j := 42 // or you can have it automatically declare the type and assign the value
	fmt.Println("j =", j)
	var k int = 10
	fmt.Println("k =", k)

	fmt.Printf("print value %v, print type %T\n", k, k)

	// if you convert int to string, the compiler will acc fetch the unicode character
	// equal to that int. See below

	k = 42
	stringK := string(k)
	fmt.Printf("the 42nd unicode character is: %v\n", stringK)

	// to convert strings correctly we shall import strconv package
	stringK = strconv.Itoa(k)
	fmt.Printf("k:%v type:%T\n", stringK, stringK) // strconv.Itoa is integer to ascii

	// values are assigned the value 0 when initialised
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	// operations
	// & AND
	// | OR
	// ^ XOR
	// &^ XAND

	// bit shifting
	// a << n shifts bits to the right which is *2^n
	// a >> n shifts bits to the left which is /2^n
	fmt.Println(4 << 3) // 4 * 2^3
	fmt.Println(4 >> 3) // 4 / 2^3

	// built in complex numbers
	comp := 1 + 1i
	// var comp complex128 = complex(1,1)
	fmt.Printf("%v, %T\n", comp, comp)
	fmt.Printf("%v, %T\n", real(comp), real(comp))
	fmt.Printf("%v, %T\n", imag(comp), imag(comp))

	// strings are aliases for bytes so
	str := "Hello World!"
	fmt.Printf("%v, %T\n", str[0], str[0])
	fmt.Printf("%v, %T\n", string(str[0]), string(str[0]))

	// strings are arrays of bytes that are converted by utf-8 and are denoted with " "
	// runes are collections of int32 that are converted by utf-32 and are denoted with ' '
	s := 'a'
	fmt.Printf("%v, %T\n", s, s)

	// constants are like variable but theyre immutable, replace var with const
	// you cant set your constants to something that will be determined at runtime
	// you can do implicit conversions with constants like so
	const a = 42
	var b float32 = 1.2
	fmt.Printf("%v, %T\n", a+b, a+b)

}
