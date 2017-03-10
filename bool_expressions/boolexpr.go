package main

import "fmt"

func main() {
	if true || false { // double pipe means "or"
		fmt.Println("This ran")
	}

	if true && false { // && means "and"
		fmt.Println("This does not run")
	}

	if !true { // ! means "not"
		fmt.Println("This did not run")
	}
}
