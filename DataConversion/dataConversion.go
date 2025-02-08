package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42
	fmt.Printf("The value is %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("The data type is %T", data)

	num1 := 123
	string1 := strconv.Itoa(num1)
	fmt.Printf("The value typr is %T", string1)
}
