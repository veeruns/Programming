package main

import (
	"fmt"
	"os"
)

func main(){
	var messageString string
	messageString = "The quick brown fox jumped over the lazy dog"

	n,err := fmt.Printf("Print stuff : %s\n",messageString)
	n=0	
        switch {
		case err != nil :
			os.Exit(1)
		case n == 0:
			fmt.Printf("No bytes output")
		case n < 26:
			fmt.Printf("Looks like it did not work  correctly %d\n",n)
		default:
			fmt.Printf("Looks like it worked correctly")

	}

		fmt.Printf("\n")

}
