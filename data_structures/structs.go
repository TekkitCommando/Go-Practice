package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type teacher struct {
	person
	canTeach bool
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := teacher{
		person:   person{"Ethan", "Smith", 19},
		canTeach: true,
	}
	p2 := teacher{
		person:   person{"John", "Doe", 20},
		canTeach: false,
	}
	fmt.Println(p1.fullName(), p1.canTeach)
	fmt.Println(p2.fullName(), p2.canTeach)
}
