package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	// Range y Close
	close(c)
	// c <- "Mensaje 3" // Error: panic: send on closed channel

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Recibido de email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Recibido de email2: ", m2)
		}
	}
}
