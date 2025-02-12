package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	fmt.Println("-------------------------------------------------------------------------------")

	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i)
	fmt.Println(f)

	fmt.Println("-------------------------------------------------------------------------------")
	const value = 47
	var j int = value
	var k float64 = value
	fmt.Println(j)
	fmt.Println(k)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	testb := b + 1
	testsmallI := smallI + 1
	testbigI := bigI + 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	fmt.Println(testb)
	fmt.Println(testsmallI)
	fmt.Println(testbigI)
}
