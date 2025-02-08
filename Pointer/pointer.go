package main

import "fmt"

func modifyValueByRef(value *int){
	*value = *value+20
}

func main (){
	var num int = 2

	var ptr *int = &num

	// fmt.Println("The valur is:", num)
	fmt.Println("Pointer containe:", ptr)
	fmt.Println("Data contains through pointer :", *ptr)

	value := 10

	modifyValueByRef(&value)
	fmt.Println("Value Contains", value)


}