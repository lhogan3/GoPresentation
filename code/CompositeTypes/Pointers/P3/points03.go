package main
import (
	"fmt"
)
// START OMIT
// Address-of(&) operators
func noPointer() string {
	return "string"
}
// You cannot return nil from a function that returns a string but you can return
// nil from a function that returns a pointer to a string!
func pointerTest() *string {
	return nil
}
func pointerTestTwo() *string {
	s := "string"                 // &"string" doesn't work
	return &s
}
func main() {
	fmt.Println(noPointer())      // prints string
	fmt.Println(pointerTest())    // prints <nil>
	// prints something like 0x40c158 (an address in memory)
	fmt.Println(pointerTestTwo())
	s := pointerTestTwo()         // now s holds an address in memory for a variable
	fmt.Println(s)                // something like 0x40c138
	sp := *s                  // now sp holds the value found at the address s holds
	fmt.Println(sp)           // string
}
// END OMIT