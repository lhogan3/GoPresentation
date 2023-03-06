package main
import "fmt"

func main(){
        fmt.Println("one")
        defer fmt.Println("three") // defer will happen before the panic
        panic("a panic happened") // the panic will happen as the last event
        fmt.Println("four") // this will never be printed
}