package main

import (
	"fmt"
	"os"
)

func main() {
	if numberBytes,theError := fmt.Printf("Hello World\n"); theError != nil {
		os.Exit(1)
	}else{
		fmt.Printf("Printed %d bytes \n",numberBytes)
	}
	
}
