package main

import (
	"fmt"
)

func main() {

        var test bool
        var bound int
        var last_prime int
	var diff int
        bound = 1000000
	for i:=1;i<bound;i++ {
		test=is_prime(i)
		if(test){
			diff = i - last_prime
		if (diff > 246 ){
			fmt.Printf("Weird stuff %d, %d,%d\n",i,last_prime,diff)
		}else{
			fmt.Printf("%d %d Diff is %d\n",i,last_prime,diff)
		}
			last_prime=i
		}
	}
       


}


func is_prime (input int) bool {
	for i:=2 ;i < (input -1); i++ {
		var modulo int
                modulo = input % i
		if(modulo == 0 ) {
			return false
                }
         }
         return true
}
