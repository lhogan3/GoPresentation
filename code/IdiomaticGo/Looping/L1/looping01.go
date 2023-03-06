// Looping - basic for-loop
package main
import "fmt"

func main() {

	total := 0  // init total
	
	// init;      condition; post statement
	for dojo := 0; dojo < 4; dojo++ {
		fmt.Println("The value of dojo -> ", dojo)
		
		total += dojo   // use total to keep track of dojo
	}

	fmt.Println("\nThe value of total is", total)
	
}
