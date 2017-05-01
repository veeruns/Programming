package main

import "fmt"

func main() {
	x := []int{8, 96, 86, 6, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 95, 17}
	var smallest int = 0
	smallest = x[0]
	for _, value := range x {
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println(smallest)
}
