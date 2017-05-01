package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64
	a = 5
	b = 10

	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	d := a - b

	fmt.Printf("%d - %d=%d\n", a, b, d)

	e := math.Pow(a, b)
	f := math.Sin(e)
	g := math.Cos(0)
	h := math.Sqrt(e)

	fmt.Printf("%f^%f=%f\n", a, b, e)
	fmt.Printf("Sine of %f is %f\n", e, f)
	fmt.Printf("Cosine of %f is %f\n", e, g)
	fmt.Printf("Sqrt of %f is %f\n", e, h)
}
