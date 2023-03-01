// Variables - Package and function level

package main

import "fmt"

// START OMIT
var palm, python, tiger bool // palm, python, and tiger are all type, "bool"

var rock, scissors int = 42, 55 // rock has a value of 42, and scissors has a value of 55

func main() {
	var i int // i is type, "int"

	// sun has a value of "eclipse", and moon has a vlue of "waxing"
	var sun, moon = "eclipse", "waxing"

	// the values of "bool" will be "false", and the value of "int" will be 0
	fmt.Println(i, palm, python, tiger)

	// "The value of the var rock is: 42"
	fmt.Println("The value of the var rock is:", rock)
	// "The value of the var scissors is: 55"
	fmt.Println("The value of the var scissors is:", scissors)

	// Concatenating strings - "Look at the moon when it is waxing."
	fmt.Println("Look at the moon when it is", moon + ".")
	fmt.Println("A total", sun, "of the Sol.") // "A total eclipse of Sol
}
// END OMIT