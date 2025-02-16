package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.1416
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculo del area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)
	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Area de un rectangulo
	baseRectangulo := 20
	alturaRectangulo := 30
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	// Area de un triangulo
	baseTriangulo := 50
	alturaTriangulo := 100
	areaTriangulo := (baseTriangulo * alturaTriangulo) / 2

	fmt.Println("Area triangulo:", areaTriangulo)

	// Area de un circulo
	const radioCirculo = 10
	areaCirculo := pi * radioCirculo * radioCirculo
	fmt.Println("Area circulo:", areaCirculo)

}
