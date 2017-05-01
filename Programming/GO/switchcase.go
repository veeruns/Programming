package main

import "fmt"

func main() {
	fmt.Println("Will select based on what you give")
	fmt.Print("Enter a number between 1 and 5 : ")
	var input float64
	fmt.Scanf("%f", &input)
	if input < 1 || input > 5 {
		fmt.Println("Dude Really ? ")
	} else {
		fmt.Println("Continuing... ")
	}

	switch input {
	case 1:
		fmt.Println("Selected = 1")
	case 2:
		fmt.Println("Selected = 2")
	case 3:
		fmt.Println("Selected = 3")
	case 4:
		fmt.Println("Selected = 4")
	case 5:
		fmt.Println("Selected = 5")
	default:
		fmt.Println("Possibly i should advice you about following rules")
	}
}
