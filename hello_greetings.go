package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// get a greeting message and print it.
	message1, err := greetings.Hello("Anton")
	fmt.Println(message1)

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Mario"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	fmt.Println(messages)

	// Request a greeting message.
	message, err := greetings.Hello("")

	//If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	fmt.Println(message)
}
