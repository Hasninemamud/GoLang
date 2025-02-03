package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// j := 0
	// for {
	// 	fmt.Printf("Counter: %d\n", j)
	// 	j += 1
	// 	if j == 11 {
	// 		break
	// 	}
	// }
	// num := []int{1, 2, 3, 4, 5}
	// for index, value := range num{
	// 	fmt.Println(index, value)
	// }

	data := "Hello, World"

	for index, char := range data{
		fmt.Printf("The index is %d and the character is %c\n", index, char)
	}

}
