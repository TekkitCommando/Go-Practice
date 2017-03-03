package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // value
	fmt.Println(&a) // memory address

	var b *int = &a
	fmt.Println(b)  // memory address
	fmt.Println(*b) // value
}
