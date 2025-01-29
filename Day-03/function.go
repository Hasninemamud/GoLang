package main

import "fmt"
func printfunc(a,b int)(int){
	return a+b
}
func main(){
	fmt.Println("Function in Go Lang")
	fmt.Println(printfunc(2, 3))
	
}