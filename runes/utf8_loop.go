package main

import "fmt"

/*
Rune is an alias for int32
UTF-8 is a 4byte coding scheme
Rune is a character which is shown by an integer
This helps me understand it but you probably don't :P
*/
func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
