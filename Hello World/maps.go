package main

import (
	"fmt"
)

// can inherit another struct but then you have to independently declare them
type Employee struct {
	employedOn string
	wentToUni  bool
}

type Doctor struct {
	Employee
	id        int
	name      string
	specialty []string
}

func main() {
	// array is a valid key type but slice is not
	// manipulating the map in one place will affect it all over the place

	// hashMap := make(map[int]string)
	var hashMap map[int]string = map[int]string{}
	fmt.Printf("%v, %v\n", hashMap, len(hashMap))

	// hashMap = map[int]string{1:"1", 2:"2"}
	hashMap[1] = "1"
	hashMap[2] = "2"

	value, ok := hashMap[3]
	fmt.Println(value, ok)

	delete(hashMap, 1)
	fmt.Println(hashMap)

	// structs can mix the datatypes
	doc1 := Doctor{
		id:   1,
		name: "Manos",
		specialty: []string{
			"drinking",
			"studying",
		},
	}

	fmt.Println(doc1)
	fmt.Println("name", doc1.name)

	// can also have a short lived struct
	// doc1 := struct{name string}{name: "Manos"}

	doc2 := doc1
	doc2.name = "Laskar"
	fmt.Println(doc1)
	fmt.Println(doc2) //copied (can be taxing if its large)

	// how the inheritence works
	doc1.employedOn = "2022-01-01"
	doc1.wentToUni = true
}
