package main

import "fmt"

const (
	EnglishHelloPrefix string = "Hello, "
	SpanishHelloPrefix string = "Hola, "
	spanish            string = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	switch language {
	case spanish:
		return SpanishHelloPrefix
	default:
		return EnglishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Sergio", "Spanish"))
}
