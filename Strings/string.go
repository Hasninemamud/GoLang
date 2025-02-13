package main

import (
	"fmt"
	"strings"
)

func main (){
	data := "apple, orange, banana"

	part := strings.Split(data, ",")
	fmt.Println(part)
	fmt.Println(len(part))

	str := "0ne, two, three, four, two"
	fmt.Println(strings.Count(str, "two"))

	str1 := "   Hello, GO! "
	removeSpace := strings.TrimSpace(str1)
	fmt.Println(removeSpace)


	str2 := "Hasnine"
	str3 := "Mamud"
	join := strings.Join([]string{str2,str3}, " ")
	fmt.Println(join)
}