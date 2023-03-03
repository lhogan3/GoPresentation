/* Alta3 Research | RZFeeser
Testing - a simple test file */

package main

import (
	"errors"
	"fmt"
	"testing"   // used for testing
	"regexp"    // regular expression library
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
// If no name was given, return an error with a message.
if name == "" {
	return name, errors.New("empty name")
}
// Create a message using a random format.
message := fmt.Sprintf("Hello there %v", name)
return message, nil
}

func TestHello(t *testing.T) {
	testName := "RZFeeser"
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