package main

import "fmt"

func arrayExample() {
	var x [58]int

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	x[42] = 777

	fmt.Println(x[42])
}

// Size is unchangeable (not dynamic) but slices are
