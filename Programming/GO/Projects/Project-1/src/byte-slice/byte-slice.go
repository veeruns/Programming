package main

import (
	"fmt"
	"os"
)

func main () {
		fileHandle,err := os.Open("test.txt")
		if(err != nil ){
			fmt.Printf("Error is %s\n",err)
			os.Exit(1)
		}
		defer fileHandle.Close()
		
		b := make([]byte,100)
		
		n,err := fileHandle.Read(b)
	
		stringVersion := string(b)

		fmt.Printf("%d : %s\n",n,stringVersion)

}
