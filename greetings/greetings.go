// declaring the package to collect related functions
package greetings

import "fmt"

// A function with capitalized letter can be called by a
// function not in the same module
func Hello(name string) string {
	// the := operator a is shortcut for declaring and
	// initialing a variable in one line (Go uses the
	// value on the right to determine the variableś type).
	// long way:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
