package kid

import "fmt"

// Hello returns a greeting for the named person.
func Hello_kid(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Hello Kids!", name)
    return message
}