package main

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	myOtherCar := car{brand: "Chevrolet", year: 2021}

	println(myCar.brand)
	println(myOtherCar.year)

	// Otra maner
	var otherCar car
	otherCar.brand = "Toyota"
	otherCar.year = 2022

	println(otherCar.brand)

}
