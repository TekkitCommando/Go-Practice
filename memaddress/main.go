package main

import "fmt"

func showMemAddress() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
}
