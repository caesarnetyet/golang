package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Hello("Julio", "English"))
}

func Hello(name string, language string) string {
	fName := strings.Trim(name, " ")
	if fName == "" {
		name = "World"
	}
	return prefixLanguage(language) + name
}

func prefixLanguage(language string) string {
	switch language {
	case "Spanish":
		return "Hola, "
	case "French":
		return "Bonjour, "
	default:
		return "Hello, "
	}
}
