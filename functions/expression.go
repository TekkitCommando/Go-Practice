package main

import "fmt"

func main() {
	/* One Way
	// This is a func expression
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	*/
	// Another way
	greet := makeGreeter()
	fmt.Println(greet())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}
