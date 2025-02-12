package main

import (
	"fmt"
)

func UpdateSlice(s []string, w string) {
	s[len(s)-1] = w
	fmt.Println(s)
}

func GrowSlice(s []string, w string) {
	s = append(s, w)
	fmt.Println(s)
}

func main() {
	s := []string{"Hello"}
	w := "World"
	UpdateSlice(s, w)
	GrowSlice(s, w)
}
