package main
import (
	"fmt"
)
// START OMIT
// With Pointers - Why we need pointers
type User struct {
	Name string
	Pet  []string
}

func (p2 *User) newPet() {
	fmt.Println(*p2, "underlying value of p2 before")
	p2.Pet = append(p2.Pet, "Fido")
	fmt.Println(*p2, "underlying value of p2 after")
}

func main() {
	// this time we'll generate a pointer
	u := User{Name: "Anna", Pet: []string{"Bailey"}}
	fmt.Println(u, "u before")
	p := &u // Let's make a pointer to u
	
	p.newPet()
	fmt.Println(u, "u after") // Does Anna have a new pet now? Yes!
}
// END OMIT