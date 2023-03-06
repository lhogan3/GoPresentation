// Calling the panic() function directly
package main
import "fmt"

func main() {
   // panic produces a quick exit
   panic("Jim, we have a problem.") // this line will be displayed

   // ...and this line will not.
   fmt.Println("You will not even see this line. The panic creates a fast fail.")
}
