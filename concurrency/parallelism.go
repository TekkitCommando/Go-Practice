package main

/*
import (
	"runtime"
	"sync"
	"fmt"
)

var wg sync.WaitGroup // Must have a wait group or main will just exit.

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // No longer needed to be included just for example
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
*/