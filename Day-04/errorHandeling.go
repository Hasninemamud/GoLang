package main

import "fmt"

func add(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error Occur")
	}
	return a / b, nil
}

func main() {
    /*"_"  blank identifier, which means it is a placeholder for 
	values that are returned by functions but are intentionally ignored"*/
	rslt, _ := add(16, 4)
	fmt.Println(rslt)
}
