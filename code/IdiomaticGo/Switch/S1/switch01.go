// switch - case and default
package main
import (
	"fmt"
	"runtime"
)
func main() {
			// init gover                
	switch gover := runtime.Version(); gover {
	// if matched "go1.19", do the code below and then STOP
	case "go1.19":
		fmt.Println("Released in 2022")
	// if matches "go1.18", do the code below then STOP
	case "go1.18":  
		fmt.Println("You are using the latest version of GoLang")
	case "go1.17":
		fmt.Println("This version of Go is fine")
	// if matches "go1.16", "go1.16.5", OR "go1.15" 
	case "go1.16", "go1.16.5", "go1.15":
		fmt.Println("You are using an older, but acceptable version of GoLang")
	default: // in all other cases run the code below
		fmt.Println("I cannot make a recommendation.")
	}
}
