package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm self executed.")
	}()
}

// Anonymous self executed func. Not very ideomatic. Will rarely be seen in Go.