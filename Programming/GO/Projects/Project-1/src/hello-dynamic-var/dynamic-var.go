package main

import (
	"fmt"
)
const (
	localScope = 43
	GlobalScope = iota * 2
   	GlobalScope2   
)
	
func main() {
	message := "The answer to life is %d\n"
 
	answer := 42


	fmt.Printf(message,answer)
	fmt.Printf(message,GlobalScope)
	fmt.Printf("%d %d\n",GlobalScope,GlobalScope2)

}
