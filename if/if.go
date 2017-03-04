package main

import "fmt"

func main() {
	if true {
		fmt.Println("This will print")
	}

	if false {
		fmt.Println("This won't print")
	}

	// Exclamation means "not"
	if !true {
		fmt.Println("This shouldn't print")
	}

	if !false {
		fmt.Println("This should print")
	}

	b := true
	d := false
	/*
		You can initialize the food var in your if statement
		but the var b is the evaluation
	*/
	if food := "Chocolate"; b {
		fmt.Println(food)
	} else if d { // Can have as many else if's as you want
		fmt.Println("this won't run because b was true")
	} else {
		fmt.Println("this won't run because b was true")
	}
}
