// Defining Methods
package main
import "fmt"
// START OMIT
type author struct {
    name      string
    books     int
}
// Method that increases the value of books by one
func (a *author) bookup() {
    a.books++
}
// Method with a receiver of author type
func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Published articles: ", a.books)
}
func main() {
    // Initializing the values of the author structure
    res := author{ name:"Liam", books:14 }
    // Calling the method
    res.show()
    // Calling bookup
    res.bookup()
    // Showing again after bookup
    res.show()
// END OMIT
}