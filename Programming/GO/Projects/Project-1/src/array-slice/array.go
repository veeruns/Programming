package main

import (
	"fmt"
)


func printer(w []string){
	for _,word := range w {
		fmt.Printf("%s\n",word)
	}
}

func main() {
	words := []string{"the","quick","brown","fox","jumped","over","the","lazy","dog"}
	printer(words[3:len(words)])
}
