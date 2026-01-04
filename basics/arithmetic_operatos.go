package basics

import (
	"fmt"
	"math"
)

func main() {
	//Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	//Overflow with sign integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//Overflow with unsign integers
	var uMaxInt uint64 = 18446744073709551615 // max value for unit64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	println(smallFloat)

}
