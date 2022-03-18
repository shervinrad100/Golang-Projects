package main

import (
	"fmt"
	"net/http"
)

func main() {
	// panic is like raising an exception in python
	// panic prints out the string you pass to it
	// i,j := 1, 0
	// ans := i/j
	// fmt.Println(ans)

	// here we shall host a web server on local host, if the port is blocked we panic

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("panics_demo.go: "))
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error()) // if port is blocked
	}

	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")        // this line will not be executed
	defer fmt.Println("end ") // this will execute before your panic
}
