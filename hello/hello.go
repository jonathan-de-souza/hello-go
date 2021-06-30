package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("using quote package")
	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Println("using greetings package")
	fmt.Println(message)
}
