package main

import (
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	age       int
}

func MakePerson(Firstname string, Lastname string, age int) Person {
	return Person{
		Firstname,
		Lastname,
		age,
	}
}

func MakePersonPointer(Firstname string, Lastname string, age int) *Person {
	return &Person{
		Firstname,
		Lastname,
		age,
	}
}

func main() {
	p := MakePerson("Chris", "Pirnack", 24)
	p2 := MakePersonPointer("Chris", "Pirnack", 24)
	fmt.Println(p)
	fmt.Println(p2)
	x := 10
	ptr := &x
	y := *ptr + 10
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(y)
	fmt.Println(x)
}
