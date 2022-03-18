package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// defer runs the statement after the function runs all of its commands but before it returns
	// just before the function is about to exit
	// defer works in last in first out so the last defer statement runs first
	// this is used to close out resources

	response, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // you close out resources straight after calling them with defer
	robots, err := ioutil.ReadAll(response.Body)
	// you may want to leave the body open in an application and work with the resource
	// but you also dont want to forget to close the resource after your function runs so you defer it
	// response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
	fmt.Println("*******************")

	// when you call defer it takes the arguments when the defer function is called not at the end
	// so in this below example youll get start not end
	a := "start"
	defer fmt.Println(a)
	a = "end"

}
