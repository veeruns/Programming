package main

import (
	"fmt"
        "os"
)

func main () {
	var s,sep string
        for index,val:= range os.Args {
	    s += sep + val
            sep = " "
            fmt.Println(index," Value is ",val)
        }
	fmt.Println (s)
}
