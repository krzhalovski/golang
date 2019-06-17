package main

import (
	"fmt"
	"strconv"
)

// Person : A simple person struct
type Person struct {
	fName string
	lName string
	age   int
	city  string
}

// Value reciever function
func (p Person) greet() string {
	p.age++
	return "Hello, " + p.fName + " " + p.lName + " " + strconv.Itoa(p.age)
}

// Pointer reciever function
func (p *Person) addAge() {
	p.age++
}

func main() {
	person := Person{"Danail", "Krzhalovski", 21, "Skopje"}

	fmt.Println(person)
	person.age++
	fmt.Println(person.greet())
	person.addAge()
	person.addAge()
	person.addAge()
	person.addAge()
	fmt.Println(person.greet())
	fmt.Println(person.greet())

}
