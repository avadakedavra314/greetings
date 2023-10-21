package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("the name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// returns random gerrting message
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!"}

	return formats[rand.Intn(len(formats))]
}
