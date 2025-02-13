package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("text.txt")
	// if err != nil{
	// 	fmt.Println("Error Occur", err)

	// }
	// defer file.Close()

	// content := "Hello World!"

	// io.WriteString(file, content)
	// // if errors != nil{
	// // 	fmt.Printf("Error Occur")
	// // 	return 
	// // }
	// fmt.Println("File successfully created")


	open, err := os.Open("text.txt")

	if err != nil{
		fmt.Println("Error Found")
		return 
	}
	defer open.Close()



}
