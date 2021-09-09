package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool{
	text = strings.ToLower(text)

	var textReverse string

	for i:= len(text) - 1; i >= 0; i--{
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	} else {
		return false
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	wasPali := isPalindromo("Ama")
	fmt.Println(wasPali)
}