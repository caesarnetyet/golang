package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Hello("Julio"))
}

func Hello(name string) string {
	fName := strings.Trim(name, " ")

	if fName == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
