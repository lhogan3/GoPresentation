// Constants - Introduction to numerical constants

package main

import "fmt"
// START OMIT
// this is a group of constants being declared at once
const (
	Code    = 585        // go will determine this is a "untyped" integer
	Big     = "New York"
	Country = "USA"
)
func oven(i string) {
	return
}
func main() {
	// display our global constants
	fmt.Println(Big)      // displays New York
	fmt.Println(Country)  // displays USA
	// Note: constants may also be lower case, they do not need to be upper case
	const Small = 3    // this is an "untyped" integer
	fmt.Println(Small)
	// if uncommented, this will fail
	// Small = 10  // reassignment is NOT allowed

	// if uncommented, this will fail
	// the value of constants needs to be known at compile time
	// therefore values returned by function calls are illegal
	// const cupcake = oven("vanilla")
}
// END OMIT