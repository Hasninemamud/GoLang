package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}
