package main

import "fmt"

import "strings"

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for _, valor := range slice {
		println(valor)
	}

	isPalindromo("anitalavalatina")
}
