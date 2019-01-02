package main

import (
	"fmt"
	"strconv"
)

// Define Person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Alternate short ways..
type PersonModified struct {
	firstName, lastname, city, gender string
	age                               int
}

//Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " , and I am " + strconv.Itoa(p.age) + " years old"
}

//Pointer reciever pass by reference
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever) //Bad example..socially incovinient :-)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	//Init Person using struct -- Preferred Approach code readability
	person1 := Person{firstName: "Joe", lastName: "Doe", city: "Boston", gender: "m", age: 25}

	//Alternate
	person2 := Person{"Samantha", "Smith", "Boston", "f", 25}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)
	fmt.Println(person1.city)

	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Thompson")
	fmt.Println(person1.greet())

	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}
