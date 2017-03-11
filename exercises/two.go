/*
Modify the first program to use a func expression
*/

package main

import "fmt"

// The answer
func two() {
	one := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(one(4))
}
