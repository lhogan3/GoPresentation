package main
import "fmt"

func printAllOperations(x int, y int) {
	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("sum=%v, divide=%v, multiply=%v \n", sum, divide, multiply)
}

func main() {
	x := 0
	y := 20
	printAllOperations(x, y)
}