package main

import "fmt"

func nestedLoop() {
	for i := 0; i <= 10; i++ {
		for j := 0; j < 10; i++ {
			fmt.Println(i, "-", j)
		}
	}
}
