package main

import (
	"fmt"
)

func main(){
	daysOfMonth := map[string]int {
		"Jan": 31,
		"Feb" : 28,
		"Mar" : 31,
		"Apr" : 30,
		"May" : 31,
		"Jun" : 30,
		"Jul" : 31,
		"Aug" : 31,
		"Sep" : 30,
		"Oct" : 31,
		"Nov" : 30,
		"Dec" : 31,
	}

	fmt.Printf("Looping through all elements \n")
	for month,days := range daysOfMonth {
		fmt.Printf("For month %s days are %d\n",month,days)
	}

	has31 := 0
	for _,days := range daysOfMonth {
		if days == 31 {
			has31 += 1

		}
	}
	fmt.Printf("Number of months with 31 days are %d\n",has31)
}
