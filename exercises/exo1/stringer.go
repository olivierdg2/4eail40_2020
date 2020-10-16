package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    fmt.Stringer
}
type Circle struct {
    r float64
  }
type Rectangle struct {
    Width  float64
    Length float64
}
func (r Rectangle) Area() float64{
    return r.Width  * r.Length
}
func (c Circle) Area() float64{
    return 3.14*math.Pow(c.r,2)
}
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle of width %f and length %f", r.Width, r.Length)
}
func (c Circle) String() string {
    return fmt.Sprintf("Circle of radius %f", c.r)
}
func main() {
    r := &Rectangle{2, 3}
    c := &Circle{5}

    shapes := []Shape{r, c}

    for _, s := range shapes {
        fmt.Println(s)
        fmt.Println("And area:",s.Area())
    }

}