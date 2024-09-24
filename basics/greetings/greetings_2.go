//greetings_2.go

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// 	return formats[...]: Retorna uma string aleatória do slice formats. A parte rand.Intn(len(formats))
// gera um número inteiro aleatório que varia de 0 até len(formats) - 1, garantindo que o índice esteja
// dentro do limite do slice.

// 	rand.Intn(len(formats)): Chama a função Intn do pacote rand, que gera um número inteiro aleatório menor
// que o valor passado (neste caso, o tamanho do slice formats).
