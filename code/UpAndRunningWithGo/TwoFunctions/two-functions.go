// Multiple functions

package main

import "fmt"

//               (param, param, param) return
func shippingCalc(x int, y int, z int) int { // the return type is int
	return x + y + z 
}

func main() {
	
	fmt.Println("Enter Height: 42")
	fmt.Println("Enter Width: 13")
	fmt.Println("Enter Length: 20")
	fmt.Println("Total: ", shippingCalc(42, 13, 20)) // print results
}
