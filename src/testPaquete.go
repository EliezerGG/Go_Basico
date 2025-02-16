package main

import (
	pk "Curso_Go/myPackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 2020
	fmt.Print(myCar)

}
