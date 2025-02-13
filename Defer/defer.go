package main 

import "fmt"

func sumnum (a,b int)(int) {
	return a+b 
}

func main (){
	fmt.Println("Hi")
	defer fmt.Printf("are you")
	defer fmt.Println(sumnum(10,12))
	fmt.Println("How")
}