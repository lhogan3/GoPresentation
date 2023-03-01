package main

import (
	"fmt"
	"math"
)
// START OMIT
func main() {
	// x and y are integers
	var earth, mars int = 3, 4
	fmt.Printf("The type of earth is %T and the value is %v\n", earth, earth)
	fmt.Printf("The type of mars is %T and the value is %v\n", mars, mars)
	
	// func float64() transforms input into float64 type
	var f float64 = math.Sqrt(float64(earth*earth + mars*mars))
	fmt.Printf("The type of f is %T and the value is %v\n", f, f)
	
	// transforn into unsigned integer
	var zebra uint = uint(f)  // we are passing by assignment
	fmt.Printf("The type of zebra is %T and the value is %v\n", zebra, zebra)

	// Inc/Dec statements
	// add one to earth
	earth++ // increase by the "untyped constant" 1
	fmt.Printf("Value increased by 1. Type is still %T and the value is now %v\n", earth, earth)

	// remove one from mars
	mars--  // decrease by the "untyped constant" 1
	fmt.Printf("The type of mars has decreased by one. The type is still %T and the value is now %v\n", mars, mars)
	// All of our operations have been passing by value (not by pointer)
}
// END OMIT