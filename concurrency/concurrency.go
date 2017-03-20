package main

/*
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Must have a wait group or main will just exit.

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
*/
