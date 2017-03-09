package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // Defers running right before the function exits, this will print "hello world"
	hello()
}
