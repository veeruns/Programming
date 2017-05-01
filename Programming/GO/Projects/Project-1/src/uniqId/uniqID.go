package main

import (
	"fmt"
)

func makeUniq(idChan chan int) {
	var id int
	id = 0
	for {
		idChan <- id
		id += 1
	}
}


func main () {
	ids := make(chan int)
	
	go makeUniq(ids)
	

	fmt.Printf("%d\n",<-ids) 
	fmt.Printf("%d\n",<-ids) 
	fmt.Printf("%d\n",<-ids) 
	fmt.Printf("%d\n",<-ids) 
}
