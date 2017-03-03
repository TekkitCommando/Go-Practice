package main

import "fmt"

func noCondition() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}
