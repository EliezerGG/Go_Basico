package main

import "fmt"

func main() {
	// Declaracion de variables
	helloMessage := "hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "USAC"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println("Mensaje:", message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
