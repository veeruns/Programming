package main

import (
	"fmt"
)

func throwUp(wordChan chan string, done chan bool) {
	words := []string{"The","Quick","Brown","Fox","Jumped","Over","The","Lazy","Dog"}

	i := 0
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
			}
	}
}


func main() {
	wordC := make(chan string)
	doneC := make(chan bool)

	go throwUp(wordC,doneC)

	for i:= 0;i < 100; i++ {
		fmt.Printf("%d: %s\n",i,<-wordC)
	}
	doneC <- true

}
