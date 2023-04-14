package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Dog struct {}
	func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

type Friend interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func Greet (f Friend) {
	f.SayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)

	// Control Structures. "IF" statments. Go's SWITCH
	var x int = 1
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}
}