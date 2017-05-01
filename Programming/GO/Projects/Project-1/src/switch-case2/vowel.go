package main

import (

	"fmt"
)

func main(){

	var msgString string
	msgString = "The Quick brown fox jumped over the lazy dog"

	vowels := 0

	consonants := 0

	space := 0

	zeds := 0
	for _,r := range msgString {
		switch r {
			case 'a','e','i','o','u' :
				vowels += 1
				fmt.Printf("Adding 1 to vowels (%d) : %c\n",vowels,r)
			case ' ':
				space += 1
			case 'z':
				zeds += 1
				fallthrough
			default:		
				consonants += 1
		}
	}
	fmt.Printf("Message String : %s\n",msgString)
	fmt.Printf("Number of vowels : %d\n",vowels)
	fmt.Printf("Number of consonats: %d\n",consonants)
	fmt.Printf("Number of zeds: %d\n",zeds)
	fmt.Printf("Number of spaces : %d\n",space)

} 
