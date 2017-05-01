package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Circle area calculation")

	fmt.Println("This is pretty sweet")

	fmt.Print("Enter Radius value : ")
	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Circle are with radius %.2f = %.2f\n", radius, area)
}
