package main

import "fmt"
func main() {
	studentGrades := make(map[string]int)

	studentGrades["Hemel"] = 100
	studentGrades["Hasnine"] = 90
	studentGrades["Mamud"] = 20
	studentGrades["Sifat"] = 180

	fmt.Println("Marks of Hemel:", studentGrades["Hemel"])
	fmt.Println("Marks of Hasnine:", studentGrades["Hasnine"])

	studentGrades["Hasnine"] = 95
	
	fmt.Println("Marks of Hasnine:", studentGrades["Hasnine"])
	fmt.Println("Marks of Sifat:", studentGrades["Sifat"])


	delete(studentGrades, "Sifat")
	fmt.Println("Marks of Sifat:", studentGrades["Sifat"])

	marks, bool := studentGrades["David"]
	if bool == false{
		fmt.Println("David not exist")
	}else{
		fmt.Println("Marks of Daivd", marks)
	}
	
	// fmt.Println("Is David Exists", bool)

}
