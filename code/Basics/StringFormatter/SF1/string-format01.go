// Introduction to String Formatting and formatting verbs

package main

import (
	"fmt"
	"os"
)
// START OMIT
func main() {
	fmt.Printf("bool: %t\n", true) // formatting booleans
	fmt.Printf("int: %d\n", 123) // use %d for standard, base-10 formatting
	fmt.Printf("bin: %b\n", 14) // binary representation
	fmt.Printf("char: %c\n", 33) // prints the character corresponding to the given integer
	fmt.Printf("hex: %x\n", 456) // hex encoding is possible with %x
	fmt.Printf("float1: %f\n", 78.9) // For basic decimal formatting
	fmt.Printf("float2: %e\n", 123400000.0) // %e and %E format the float in (slightly different versions of) scientific notation
	fmt.Printf("float3: %E\n", 123400000.0) // basic string printing use %s
	fmt.Printf("str1: %s\n", "\"string\"") // double-quote strings as in Go source, use %q 
	fmt.Printf("str2: %q\n", "\"string\"") // %x renders the string in base-16
	fmt.Printf("str3: %x\n", "hex this") // with two output characters per byte of input
	// specify the width of an integer, use a number after the % in the verb
	// By default the result will be right-justified and padded with spaces
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // specify the width of printed floats
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // left-justify, use the - flag
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // left-justify use the - flag as with numbers
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("sprintf: a %s", "string") // Sprintf formats and returns a string without printing it anywhere
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // format+print to io.Writers other than os.Stdout using Fprintf.
}
// END OMIT