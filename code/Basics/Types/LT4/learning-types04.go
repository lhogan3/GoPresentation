// Type transformations - String to integer

package main

import (
	"fmt"
	"strconv"    // built in package for string conversions
)
// START OMIT
func main() {

	// create a string via inference
	s := "33"
	
	// string to int using imported "strconv" package
	i, err := strconv.Atoi(s)
	
	// avoid nesting and check error immediately as less nesting means less load
	// on the compiler
	if err != nil {
		panic(err)
	}
	
	// the type of s followed by the value of s
	fmt.Printf("The type of s is %T and the value is %v\n", s, s)
	// the type of i followed by the value of i
	fmt.Printf("The type of i is %T and the value is %v\n", i, i)
			
	// transform an integer into a string
	st := strconv.Itoa(42)   // We have transformed 42 into a string
	// the type of st followed by the value of st    
	fmt.Printf("The type of st is %T and the value is %v\n", st, st)    
}
// END OMIT