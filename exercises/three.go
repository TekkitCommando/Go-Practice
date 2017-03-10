/*
Write a function with one variadic parameter that finds
the greatest number in a list of numbers.
*/
package main

import "fmt"

// The answer
func greatestNum(n ...int) {
	var greatest int
	for _, v := range n {
		if v > greatest {
			greatest = v
		}
	}
	fmt.Println(greatest)
}

func main() {
	greatestNum(1, 10, 7, 9, 12, 27)
}
