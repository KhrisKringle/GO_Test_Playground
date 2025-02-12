package main

import (
	"fmt"
	"math/rand"
)

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
	evenVals2 := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals2 {
		if i == 0 {
			continue
		}
		if i == len(evenVals2)-1 {
			break
		}
		fmt.Println(i, v)
	}
	/*loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}*/

	x := []int{}
	for i := 0; i < 100; i++ {
		n := rand.Intn(10)
		x := append(x, n)
		fmt.Println(x)
		for i := 0; i < len(x); i++ {
			test := x[i]
			switch {

			case test%2 == 0 && test%3 == 0:
				fmt.Println("Six")

			case test%2 == 0:
				fmt.Println("Two")

			case test%3 == 0:
				fmt.Println("Three")

			default:
				fmt.Println("Never Mind")
			}
		}

	}

}
