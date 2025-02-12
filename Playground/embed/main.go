package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
)

//go:embed english_right.txt
var e_right string

//go:embed german_right.txt
var g_right string

func main() {
	args := os.Args
	if len(args) > 2 {
		errors.New("Too many arguments")
	} else {
		if args[1] == "english" {
			fmt.Println(e_right)
		}

		if args[1] == "german" {
			fmt.Println(g_right)
		}
	}

}
