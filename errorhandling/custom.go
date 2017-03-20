package main

import (
	"log"
	"fmt"
)

type ErrNegSqrt struct {
	lat, long string
	err error
}

func (n *ErrNegSqrt) Error() string {
	return fmt.Sprintf("What?! A math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Math error: square root of negative number: %v", f)
		return 0, &ErrNegSqrt{"50.02", "99.50", nme}
	}
	// implementation
	return 42, nil
}
