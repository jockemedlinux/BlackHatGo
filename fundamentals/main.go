//@jockemedlinux
package main

import (
	"fmt"
	"time"
	"errors" //not mentioned in book. Error handling section wont work without it.
	// "encoding/json" Not mentioned in book. json handling structured data not working properly. Need to fix.
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

type Foo struct {
	Bar string
	Baz string
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func Greet (f Friend) {
	f.SayHello()
}


func f() {
	fmt.Println("f function here")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

type MyError string
func (e MyError) Error() string {
	return string(e)
}

func foo() error {
	return errors.New("Some Error Occurred")
}

func main() {
		fmt.Println("Hello, Black Hat Gophers!\n")
		
		//primitive data-types
		var x = "Hello World"
		z := int(42)
		fmt.Println(x)
		fmt.Println(z)
		fmt.Println("\n")

		// slices and maps
		var s = make([]string, 0)
		var m = make(map[string]string)
		s = append(s, "some string")
		m["some key"] = "some value"
		fmt.Println(s)
		fmt.Println(m)

		// Pointers, Strucs, and Interfaces
		fmt.Println("\n")
		var count = int(42)
		ptr := &count
		fmt.Println(*ptr)
		*ptr = 100
		fmt.Println(count)

		var guy = new(Person)
		guy.Name = "Dave"
		Greet(guy)
		var dog = new(Dog)
		Greet(dog)
	

		// Control Structures. "IF" statments. Go's SWITCH
		fmt.Println("\n")
		var u int = 1
		if u == 1 {
			fmt.Println("X is equal to 1")
		} else {
			fmt.Println("X is not equal to 1")
		}

		
		// SWITCHES
		fmt.Println("\n")
		var l = "foo"
		switch l {
		case "foo":
			fmt.Println("Found foo")
		case "bar":
			fmt.Println("Found bar")
		default:
			fmt.Println("Default case")
		}

		// For loops
		fmt.Println("\n")		
		for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
		fmt.Println("\n")
		nums := []int{2,4,6,8}
		for idx, val := range nums {
		fmt.Println(idx, val)
	}

	// Concurrency
		fmt.Println("\n")
		go f()
		time.Sleep(1 * time.Second)
		fmt.Println("Continued main function after f function")

	// Concurrency #2
		fmt.Println("\n")
		c := make(chan int)
		go strlen("Salutations", c)
		go strlen("World", c)
		b, d := <-c, <-c
		fmt.Println(b, d, b+d)

	// Error handling
		if err := foo();err != nil {
			//handle the error
		}


	// Structured Data Handling. Not working properly.
	//	fmt.Println("\n")
	//	f := Foo{"Joe Junior", "Hello Shabado"}
	//	b, _ := json.Marshal(f)
	//	fmt.Println(string(b))
	//	json.Unmarshal(b, &f)


}