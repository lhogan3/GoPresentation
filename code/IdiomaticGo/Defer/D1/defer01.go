// Defer - cleanup operations on file creation
package main
import (
	"fmt"
	"os"
)
// START OMIT
// expects input, "p" to be a string will return os.File object
func createFile(p string) *os.File {
	fmt.Println("Creating file...")
	f, err := os.Create(p)
	if err != nil { panic(err) }
	return f
}
// exects input, f, to be an os.File object, returns nothing
func writeFile(f *os.File) {
	fmt.Println("Writing employee list to file...")
	fmt.Fprintln(f, "Cloud Strife\nAerith Gainsborough\nRed XIII\nVincent Valentine")
}
func closeFile(f *os.File) {
	fmt.Println("Closing file...")
	err := f.Close()
	if err != nil { fmt.Fprintf(os.Stderr, "error: %v\n", err) }
}
func main() {
	// create file object 
	f := createFile("/tmp/employeeList.txt")
	// after getting back a file object, we defer closing closeFile()
	defer closeFile(f)   // closeFile() will be executed at end of the enclosing
	writeFile(f)		// function-- in this case, main()
}
// END OMIT