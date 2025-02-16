package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func doubleReturn(a int) (int, int) {
	return a, a * 2
}

func returnValue(a int) int {
	return a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1:", value1, "value2:", value2)

}
