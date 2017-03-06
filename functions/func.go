package main

import "fmt"

func funcCall() {
	fmt.Println(greet("Ethan", "Smith"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

// Greet is declared with a parameter
// when calling greet, pass in an argument
