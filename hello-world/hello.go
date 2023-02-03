package main

import "fmt"

const (
	EnglishHelloPrefix = "Hello, "
	SpanishHelloPrefix = "Hola, "
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s", EnglishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Sergio"))
}