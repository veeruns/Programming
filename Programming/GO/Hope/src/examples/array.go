package main

import (
	"fmt"
)

func main() {
	var x [6]float64
	x[0] = 54
	x[1] = 566
	x[2] = 510
	x[3] = 155
	x[4] = 25
	x[5] = 98

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	var avg float64 = 0
	avg = total / float64(len(x))

	fmt.Printf("Total is %f, Average is %f\n", total, avg)
}
