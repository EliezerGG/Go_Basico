package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y soy de la marca %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)
}
