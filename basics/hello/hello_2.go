//hello_2.go

package main

import (
	"fmt"

	"github.com/Marcelo-Magal/go-studies/basics/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Marcelo")
	fmt.Println(message)
}
