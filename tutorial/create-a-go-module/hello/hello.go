package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including the log entry
	// prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Jiseok", "Kyoseok", "Haseon"}
	// Request a greeting message.
	// message, err := greetings.Hello("Jiseok")
	messages, err := greetings.Hellos(names)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
