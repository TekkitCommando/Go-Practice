package main

import "fmt"

func mapExample() {
	m := make(map[string]int) // or m := map[string]int{"key": value}
	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")
	fmt.Println("map:", m)

	v, ok := m["k2"]
	fmt.Println("ok?:", ok, v)
}
