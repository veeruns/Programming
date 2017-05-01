package main

import (
	"fmt"
)

func Printer(msg string) error {
	_,err := fmt.Printf("Message is : %s\n",msg)
	return err
}

func main(){
	Printer("Hello World")
}
