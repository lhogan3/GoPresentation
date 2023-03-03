package main
import "fmt"

func printAllOperations(x int, y int) {
	defer func() {
		// defer function to escape the panic when y/x runs
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
			fmt.Println("Proceeding to alternative flow skipping division.")
			printOperationsSkipDivide(x, y)
		}
	}()
	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("sum=%v, divide=%v, multiply=%v \n", sum, divide, multiply)
}
func printOperationsSkipDivide(x int, y int) {
	sum, multiply := x+y, y*x
	fmt.Printf("sum=%v, multiply=%v \n", sum, multiply)
}

func main() {
	x := 0
	y := 20
	printAllOperations(x, y)
}