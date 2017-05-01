package main

import (
	"fmt"
)

func emit(ch chan string) {
	words := []string{"The","Quick","Brown","Fox","Jumped","Over","the","Lazy","dog"}
	for _,word := range words {
		ch <- word
	}
	close(ch)
}

func main() {
	wordChan := make(chan string)
	
	go emit(wordChan)
	for word := range wordChan {		
		fmt.Printf("%s ",word)
	}
	fmt.Printf("\n")
}
