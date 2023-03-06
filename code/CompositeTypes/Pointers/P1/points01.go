// The pointers of a type are initialized using the address-of(&) operator on a 
// subject of that specific type.

package main
import (
   "fmt"
)

func main() {
   var q int = 42
   // declare the pointer, "p is a pointer to a var that will be an int type"
   var p *int
   p = &q          // initialize the pointer
   fmt.Println(p)  // memory address (0x20e010)
}
