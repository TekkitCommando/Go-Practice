package scope

import "fmt"

var x int = 42 // Package level scope (Accessible only in its own package)

func main() {
	fmt.Println(x)
	foo()
	y := 17 // Block level scope (Accessible only down to the curly brace)
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
