package hello

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! We finally met", name)
	return message
}

func HelloMessage(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hello message, %v. Welcome!", name)
	return message
}

