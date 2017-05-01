package main

import (
	"fmt"
	"os"
)


func createFile(msg string) (int, error) {
		f,err := os.Create("Hello-World-File.txt")
		var i int
		i=0
		if err != nil {
			return i,err
		}
		defer f.Close()
		numberBytes,errorWrite := f.WriteString(msg)
		if errorWrite != nil {
			fmt.Printf("Error while writing file")
		}else{
			fmt.Printf("Done writing to file\n")
		}		
			return numberBytes,errorWrite
}

func main() {
		n,err := createFile("The quick brown fox jumped over the lazy dog")
		fmt.Printf("%d and %s",n,err)


}
