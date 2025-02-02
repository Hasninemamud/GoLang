package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter a number 1-7: ")
	fmt.Scan(&day)

	switch day{
	case 1:
		fmt.Print("Saturday")
	case 2:
		fmt.Print("Sunday")
	case 3:
		fmt.Print("Monday")
	case 4:
		fmt.Print("Tuesday")
	case 5:
		fmt.Print("Wednesday")
	case 6:
		fmt.Print("Thursday")
	case 7:
		fmt.Print("Friday")
	}


}
