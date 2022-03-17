package main

import (
	"fmt"
	"reflect"
)

// structs work sort of like a postgres db
// you can add tags which are constraints on the fields

type Employee struct {
	id   int
	name string `required,max:"100"` // spaces are not allowed in tags, separate with commaß
}

type Doctor struct {
	Employee
	specialty []string
}

func main() {

	doc1 := Doctor{
		Employee: Employee{
			id:   1,
			name: "Manos",
		},
		specialty: []string{
			"drinking",
			"working",
		},
	}

	fmt.Println(doc1)

	// to show the tags you use the reflect package
	t := reflect.TypeOf(Doctor{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
