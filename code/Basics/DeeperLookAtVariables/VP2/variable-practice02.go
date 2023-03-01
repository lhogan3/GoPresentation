package main

import (
	"fmt"
)
// START OMIT
// this variable is "global" for the "main" package
// because it begins with a lowercase letter
var zzz string = "package-level (global) variable\n"
func main() {
	// display our "global" variable
	fmt.Printf(zzz)
	
	// This would be "long hand" declaration
	var mug string = "filled with tea\n"
	fmt.Printf(mug)
	
	// Either type or expression may be omitted (not both)
	// These are all examples of type being omitted        
	var a = "Hello"
	var b = 23
	var c = true
	var d = 2.3
	
	// var a = "Goodbye"
	a = "Goodbye"

	// Prints: The type of a, b, c, d is: string, int, bool and float64 respectively
	fmt.Printf("The type of a, b, c, d is: %T, %T, %T and %T.", a, b, c, d)
}
// END OMIT
