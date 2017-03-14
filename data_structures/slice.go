package main

import "fmt"

func sliceExample() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}

func sliceExample2() {
	/* Doubles when capacity is filled. If you don't specify the 3rd parameter
	the length becomes both the length and the capacity */
	mySlice := make([]int, 0, 5)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}

func differentWays() {
	example := []string{} // Shorthand, needs an append
	fmt.Println(example)
	var example2 []string // Var, needs an append, defaults to nil
	fmt.Println(example2)
	example3 := make([]string, 20) // Make, ideomatic way
	fmt.Println(example3)
}

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Println(greeting)
}
