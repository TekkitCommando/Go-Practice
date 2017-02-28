package main

import "fmt"

func main() {
	name := "Ethan"
	length, _ := fmt.Print(name + "\n") // length can be used without also using err because of blank identifier
	fmt.Print(length)
}
