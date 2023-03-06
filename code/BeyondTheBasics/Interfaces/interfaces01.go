// Getting at interfaces
package main
import (
  "fmt"
  "math"
)
// START OMIT
type geometry interface {
  area() float64   // needs an area method that returns a float64
}
type rect struct {
  width, height float64
}
type circle struct {
  radius float64
}
func (r rect) area() float64 {
  return r.width * r.height
}
func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}
func measure(g geometry) { // a generic "measure" function taking advantage of the
  fmt.Println(g)           // interface "geometry"
  fmt.Println(g.area())
}
func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}
  measure(r)    // run measure against rect
  measure(c)    // run measure against circle
  // END OMIT
}