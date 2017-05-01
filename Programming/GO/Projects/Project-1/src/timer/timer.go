package main

import (
	"fmt"
	"time"
)

func throwUp(wordChan chan string, done chan bool) {
	words := []string{"The","Quick","Brown","Fox","Jumped","Over","The","Lazy","Dog"}

	defer close(wordChan)

	i := 0

	tC := time.NewTimer(3 * time.Second)
	for {
		select {
			case wordChan <- words[i]:
				i += 1
				if i == len(words) {
					i = 0
				}
			case <- done: 
				close(done)
				fmt.Printf("Closed channel done, we are done\n")
				return
			case <- tC.C:
				fmt.Printf("Timer returned\n")
				return
			
			}
	}
}


func main() {
	wordC := make(chan string)
	doneC := make(chan bool)

	go throwUp(wordC,doneC)

	for word := range wordC {
		fmt.Printf(" %s\n",word)
	}

}
