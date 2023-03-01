package main

import "fmt"
// START OMIT
// In Go, variables are explicitly declared and used by the compiler to e.g. check 
// type-correctness of function calls
func main() {
    // variables declared inside of main() are at the "function" level
    // var name type = expression
    var name string = "liam"
    
    // `var` declares 1 or more variables. here we let the compiler infer the type
    var a = "initial"

    // You can declare multiple variables at once.
    var b, c int = 1, 2

    // Go will infer the type of initialized variables.
    var d = true

    // The `:=` syntax is shorthand for declaring and initializing a variable, 
    // e.g. for `var f string = "apple"` in this case.
	// (Also called the walrus operator which is an incredible name IMO)
    f := "apple"
    // END OMIT
    fmt.Println(name)
    fmt.Println(a)
    fmt.Println(b, c)
    fmt.Println(d)
    fmt.Println(f)
}