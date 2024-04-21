package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

func (h Human) String() string {
	return fmt.Sprintf("human named %s %s of age %d living in %s", h.Firstname, h.Lastname, h.Age, h.Country)
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
	Walk()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
}

func (c Cat) Walk() {
	fmt.Printf("The %s has given affection to Human named %s\n", c.Name)
}

func (d Dog) ReceiveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	// animal.ReceiveAffection(human)
}

func main() {
	// Create the Human
	// var john Human
	// john.Firstname = "John"

	// Create a Cat
	var c Cat
	c.Name = "Maru"

	// then a dog
	var d Dog
	d.Name = "Medor"

	// Pet(c, john)
	// Pet(d, john)
	var john Human
	john.Firstname = "John"
	john.Lastname = "Doe"
	john.Country = "USA"
	john.Age = 45

	// fmt.Println(john)

	var x interface{}
	var num int = 10
	var str string = "hello"
	var fl float64 = 5. / 2

	x = num
	findType(x)
	x = str
	findType(x)
	x = fl
	findType(x)

}

func findType(x interface{}) {
	switch x := x.(type) {
	case int:
		fmt.Printf("this is int: %d\n", x*2)
	case string:
		fmt.Printf("this is string: %s\n", strings.ToUpper(x))
	case float64:
		fmt.Printf("this is float64: %.2f\n", x*10)
	default:
		fmt.Println("unknown type")
	}
}
