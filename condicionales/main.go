package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es uno")
	} else {
		fmt.Println("No es uno")
	}

	// with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es uno y dos")
	}

	// with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es cero o dos")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Value", value)
	}
}
