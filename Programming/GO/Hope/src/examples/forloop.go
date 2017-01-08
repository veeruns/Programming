package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Hello %d (Even) \n", i)
		} else {
			fmt.Printf("Hello %d (Odd) \n", i)
		}
	}
}
