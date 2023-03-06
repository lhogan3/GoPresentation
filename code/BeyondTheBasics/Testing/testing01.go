// Testing - a simple test file
package main
import (
	"errors"
	"fmt"
	"testing"   // used for testing
	"regexp"    // regular expression library
)

// START OMIT
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
if name == "" { 
	return name, errors.New("empty name") // If no name given, return an error.
}
message := fmt.Sprintf("Hello there %v", name) // Creates message using random format.
	return message, nil
}
func TestHello(t *testing.T) {
	testName := "Liam"
	desiredResult := regexp.MustCompile("Hello there "+testName)
	result, err := Hello("RZFeeser")
	if !desiredResult.MatchString(result) || err != nil {
		t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
	}
}
func TestHelloEmptyString(t *testing.T) {
	result, err := Hello("")
	if result != "" || err == nil {
		t.Fatalf("Failed to produce valid error for empty string. Got, %v", result)
	}
}
// END OMIT