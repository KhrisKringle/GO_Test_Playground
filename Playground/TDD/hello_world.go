package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "

	french  = "French"
	spanish = "Spanish"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingsPrefix(language) + name
}

func greetingsPrefix(lang string) (prefix string) {
	prefix = englishHello

	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	}

	return prefix
}

var data = `
a: Hello
b:
  c: [1, 2]
`

type Info struct {
	A string
	B struct {
		C []int
	}
}

func main() {
	fmt.Println(Hello("Chris", french))

	info := Info{}

	err := yaml.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", info)
}
