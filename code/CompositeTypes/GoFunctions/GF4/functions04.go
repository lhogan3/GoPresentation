// Function literals - Anonymous
package main
import (
	"fmt"
	"sort"
)
// START OMIT
func main() {
	// people is a slice of structures we are defining "anonymously", which means
	// to instantiate the instance immediately after declaring the type
	people := []struct {
		Host string
		Mac  string
		Ram  int
		Cpu  int
	}{
		{"merlin", "2e549138a9e3", 1024, 1},
		{"prospero", "3c539188c9e3", 2048, 2},
		{"nestor", "1b166127a9e3", 2048, 1},
	}
	// An anonymous function is a function which doesn’t contain any name.
	// It is useful when you want to create an inline function.                                         
	sort.Slice(people, func(i, j int) bool { return people[i].Host < people[j].Host }) 
	fmt.Println("By Hostname:", people) // sort in ascending order
	sort.Slice(people, func(i, j int) bool { return people[i].Ram < people[j].Ram })
	fmt.Println("By RAM:", people)
	// assigning a nameless function to a variable is also possible
	algo := func(i, j int) bool { return people[i].Mac > people[j].Mac }
	sort.Slice(people, algo)
	fmt.Println("Reverse sort by MAC:", people)
}
// END OMIT