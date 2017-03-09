package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Ethan], didn't have to put the address because it is a reference type
}

func changeMe(z []string) {
	z[0] = "Ethan"
	fmt.Println(z) // [Ethan]
}
