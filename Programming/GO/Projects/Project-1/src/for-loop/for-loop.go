package main

import(
	"fmt"
)

func main() {

	atoz := "The quick brown fox jumped over the lazy dog"

	for i,r := range atoz {
		
		fmt.Printf("Index : %d and Value %c\n",i,r)
	}

	
}
