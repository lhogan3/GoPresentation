// Simple if statements
package main
import ( 
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ensure we produce random values, we need to change rand.Seed() from something
	// other than rand.Seed(1), which is the default value
	rand.Seed(time.Now().UnixNano())
	// simple if statement    
	if 10%5 == 0 {                   // if 10 mod 5 is 0
		fmt.Println("10 is divisible by 5")            
	}
	// declare variable; and if this condition is true...
	// v := rand.Intn(101)
	// if v > 60 { code }
	if v := rand.Intn(101); v > 60 {
		fmt.Println("You tested in the top 60%! Your score was", v)
	}
	fmt.Println("We want to thank everyone for taking the test.")
}