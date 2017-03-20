package main

import (
	"os"
	"fmt"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("An error occured:", err) // or log.Println or log.Fatalln or panic
	}
}
