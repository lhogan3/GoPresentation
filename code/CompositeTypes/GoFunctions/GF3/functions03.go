package main
import (
	"errors" // preview of using errors
	"fmt"
)
// START OMIT
/* Errors are "just" a type that can be returned by functions. This simple approach
might be a bit different than other languages you've used. Errors can be constructed 
on the fly using Go’s built-in errors or fmt packages.In this example we look at 
using, errors.New(). errors.New constructs a basic error value with given message */

func rollchar(firstName string, lastName string) (string, error) {
	// generate an error
	if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
		return lastName, errors.New("Vegetables are not last names for heroes.")
	}
	// desirable response
	return firstName + " the " + lastName, nil
}

func main() {
	fmt.Println("Welcome to the Character Generator")
	playerChar, err := rollchar("Gandalf", "Turnip")
	if err != nil {
		fmt.Println("Error while spawning your requested character.")
	} else {
		fmt.Println(playerChar, "has been generated.")
	}
}
// END OMIT