package main
import (
	"fmt"
)
// START OMIT
// switch - using fallthrough to continue code below the case.
// In the real world, fallthrough could be use to mimic pipeline behavior.
// Start on a case like, "build" or "test" and fallthrough to the additional steps.
// our use of fallthrough creates the following pattern, A -> B -> C -> D -> E
func main() {
	nextstop := "Batcave"
	fmt.Println("Stops ahead of us:")
	switch nextstop {
		case "Superman":
			fmt.Println("Superman's Fortress of Solitude")
			fallthrough
		case "Batcave":
			fmt.Println("Batcave")
			fallthrough
		case "Strange":
			fmt.Println("Dr. Strange's Sanctum Sanctorum")
			fallthrough
		case "Justice":
			fmt.Println("Justice League's Hall of Justice")
			fallthrough
		case "Xavier":
			fmt.Println("Xavier's School for Gifted Youngsters")
		}
}
// END OMIT