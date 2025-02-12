package main

import (
	"fmt"
	"os"
)

func filelen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileinfo.Size()), nil
}

func main() {
	length, err := filelen("help.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File Length:", length, `bytes`)
	}
}
