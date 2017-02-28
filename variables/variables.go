package variables

import "fmt"

func main() {
	// Shorthand notation, this is the first acceptable way to do variables
	name := "Ethan"
	age := 19
	money := 19.99
	isTwenty := false

	// Var notation, this is the other way to do variables. Using more then just these two can affect readablility
	var varName string = "Ethan"
	var varAge int = 19
	var varMoney float64 = 19.99 // There is also float32
	var varIsTwenty bool = false

	fmt.Printf("%v \n", name)
	fmt.Printf("%v \n", age)
	fmt.Printf("%v \n", money)
	fmt.Printf("%v \n", isTwenty)
	fmt.Printf("%T \n", varName)
	fmt.Printf("%T \n", varAge)
	fmt.Printf("%T \n", varMoney)
	fmt.Printf("%T \n", varIsTwenty)
}
