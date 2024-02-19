package main

import (
	"fmt"
)

const englishPrefixHello = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return englishPrefixHello + name
}

func main() {
	fmt.Println(Hello("Jordan", "Spanish"))
}
