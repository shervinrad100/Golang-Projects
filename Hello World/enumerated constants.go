package main

import (
	"fmt"
)

// iota is a special keyword
const (
	errSpecialist = iota
	// _ = iota
	catSpecialist
	dogSpecialist
	// ...
)

// example 2
const (
	_  = iota
	KB = 1 << (10 * iota) // times by 2^(10*1)
	MB                    // times by 2^(10*2)
	GB                    // times by 2^(10*3)
	TB                    // times by 2^(10*4)
	PB                    // times by 2^(10*5)
)

// example 3
const (
	isAdmin = 1 << iota
	isHQ
	canSeeFinancials

	isAsia
	isAfrica
	isEurope
	isNorthAmerica
	isSouthAmerica
)

func main() {
	fmt.Println("enumerated constants in order")
	fmt.Println("err..", errSpecialist)
	fmt.Println("cat..", catSpecialist)
	fmt.Println("dog..", dogSpecialist)

	// since the first one is 0 if you only initialise the var then it will get the value of 0
	// so to avoid false positives you assign the first enumerated const to an error term
	// you can also assign it to a throw away variable but then we can't call the 0th term
	var mySpecialist int
	fmt.Println(mySpecialist == catSpecialist)

	var mySpecialistType = catSpecialist
	fmt.Println(mySpecialistType == catSpecialist)
	fmt.Println(mySpecialistType == dogSpecialist)

	// you can also offset the start with
	// _ = iota + 10 evaluates to -> 11
	// const1 evaluates to -> 12 and so on

	// example 2
	var fileSize = 10000000000. // Bytes
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// example 3
	// you can store user permissions in bytes
	var roles byte = isAdmin | canSeeFinancials | isEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is Asia? %v\n", isAsia&roles == isAsia)

}
