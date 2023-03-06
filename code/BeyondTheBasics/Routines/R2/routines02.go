// goroutine - with waitgroups
package main
import (
		"fmt"
		"time"
		"sync"
)
// START OMIT
func countDown(s int) {
	defer wg.Done()
	for i := s; i > 0; i-- {
			fmt.Println(i)
			time.Sleep(time.Second)   // wait one second
	}
}

// WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup
func main() {
	fmt.Println("NASA launch sequence starting:")
	wg.Add(1) // the waitGroup has "1" in the queue
	go countDown(10)
	fmt.Println("Launch sequence starting:")
	time.Sleep(time.Second)
	fmt.Println("Hey wait a second...")
	time.Sleep(time.Second)
	fmt.Println("I forgot my wallet!")
	wg.Wait() // wait for the waitGroup queue to reach 0
	fmt.Println("TAKE OFF!")
}
// END OMIT