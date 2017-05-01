package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["val"] = 10
	fmt.Println(x["key"])
	fmt.Println(x["val"])
}
