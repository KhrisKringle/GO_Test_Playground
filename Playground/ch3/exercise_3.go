package main

import (
	"fmt"
)

type Employee struct {
	firstname string
	lastname  string
	id        int
}

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	greet1 := greetings[:2]
	greet2 := greetings[1:4]
	greet3 := greetings[3:]
	fmt.Println(greet1)
	fmt.Println(greet2)
	fmt.Println(greet3)
	var chris Employee
	chris.firstname = "Chris"
	chris.lastname = "Pirnack"
	chris.id = 1
	tom := Employee{"Tom", "Monster", 2}
	marice := Employee{firstname: "Marcie", lastname: "Moncia", id: 3}
	fmt.Println(chris)
	fmt.Println(tom)
	fmt.Println(marice)
}
