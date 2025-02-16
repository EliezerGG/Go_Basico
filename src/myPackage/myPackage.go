package myPackage

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate Car with private access
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Print a message
func PrintMessage() {
	println("Hello from myPackage")
}
