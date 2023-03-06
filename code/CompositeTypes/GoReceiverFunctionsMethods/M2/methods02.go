/* Introduction to working with methods. The basic pattern of a method is as follows:
   func(receiver_name Type) method_name(parameter_list)(return_type){
      // Code
   }
*/
package main
import "fmt"
// START OMIT
type astro struct {
    name     string
    mission  string
    isneeded bool
}
func (a astro) customhonorific(honor string) string {
        return honor + a.name
}
func (a *astro) suitup() {
        a.isneeded = true
}
func main() {
    ast1 := astro{"Megan McArthur", "ISS", false}
    fmt.Println(ast1)
    fmt.Println(ast1.customhonorific("Awesome Space Hero "))
    ast1_pointer := &ast1
    fmt.Println(ast1_pointer.customhonorific("Awesome Space Hero ")) // methods move with pointers
    // try to MUTATE the data within our structure
    fmt.Println("False to True")
    fmt.Println(ast1)  // the ast1 struct has "isneeded: false"
    ast1.suitup()      // modifies the ast1 struct to "isneeded: true"
    fmt.Println(ast1)  // proof that the ast1 struct was modified
}
// END OMIT