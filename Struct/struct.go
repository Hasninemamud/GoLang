package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct{
	Email string
	Phone string

}

type Address struct{
	House int
	Area string 
	State string
}

type Employee struct{
	Persona_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// var hemel Person
	// fmt.Println(hemel)
	// hemel.FirstName = "Hasnine"
	// hemel.LastName = "Mamud"
	// hemel.Age = 24
	// fmt.Println(hemel)

	// person1 := Person{
	// 	FirstName: "Hemel",
	// 	LastName: "Mamud",
	// 	Age: 24,
	// }
	// fmt.Println("Person1:", person1)
	var employee1 Employee
	employee1.Persona_Details = Person{
		FirstName: "Hasnine",
		LastName: "Mamud",
		Age: 24,
	}
	employee1.Person_Contact.Email = "hasnine@gmail.com"
	employee1.Person_Contact.Phone = "0199882928839"

	employee1.Person_Address = Address{
		House: 12,
		Area: "Bhola",
		State: "Barisal",

	}
	fmt.Println("Employee 1:", employee1)

}
