/*
Modify the first program to use a func expression
*/

package main

import "fmt"

// The answer
func two() {
	one := func(n int) (int, bool) {
		if n%2 == 0 {
			return n / 2, true
		}
		return n / 2, false
	}

	fmt.Println(one(4))
}
