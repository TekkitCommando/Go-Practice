package main

import "fmt"

/*
No default fallthrough so you don't need break
*/
func main() {
	switch "Ethan" {
	case "Bob":
		fmt.Println("Hello Bob!")
	case "Ethan":
		fmt.Println("Hello Ethan!")
	case "Kevin":
		fmt.Println("Hello Kevin!")
	default:
		fmt.Println("I don't know your name")
	}
}

// Will continue down cases till it finds something true or is default
func switchNoCondition() {
	switch {
	case len("Ethan") == 2:
		fmt.Println("Hello")
	default:
		fmt.Println("This is default")
	}
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}
