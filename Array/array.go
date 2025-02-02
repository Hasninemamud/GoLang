package main

import "fmt"

func main() {
	var name = []string{"Jemel", "Hemel", "sifat"}
	for i := 0; i <= len(name)-1; i++ {
		fmt.Println(name[i])
	}
	// fmt.Println(name)
	// fmt.Println(len(name))
}
