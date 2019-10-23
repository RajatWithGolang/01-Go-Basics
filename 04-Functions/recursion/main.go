package main

import "fmt"

func fact(n int) int{
	
	if n > 1{
		return n * fact(n-1)
		 	}
	return 1
}

func main(){

	fact := fact(8)
	fmt.Println(fact)
}

