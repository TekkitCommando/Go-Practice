package packagescope

import "fmt"

var x int = 42 // Package level scope (Accessible only in its own package)

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
