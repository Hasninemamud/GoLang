package main

import "fmt"

func main(){
	// num := [] int{1,2,3,4,5}
	// fmt.Println("Numbers are:", num)
	// num = append(num, 6,7,8)
	// fmt.Print("Numbers are:", num)

	num := make([]int, 3, 5)
	fmt.Println("Slice: ", num)
	num = append(num, 10)
	num = append(num, 20)
	num = append(num, 10)
	fmt.Println("Slice: ", num)
	fmt.Println("Length: ", len(num))
	fmt.Println("Capacity: ", cap(num))
}