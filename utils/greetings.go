package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func RandomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	fmt.Fprint(log.Writer(), "RandomFormat() called\n")	
	return formats[rand.Intn(len(formats))]
}
