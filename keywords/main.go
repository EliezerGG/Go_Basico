package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("Es dos")
			continue
		}
		if i == 8 {
			fmt.Println("Break")
			break
		}
		fmt.Println(i)
	}

}
