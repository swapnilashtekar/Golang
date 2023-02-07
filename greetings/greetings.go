package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}

	// If name was received, return a value that embeds the name
	// in a greeting message.

	// returns a greeting taht embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
