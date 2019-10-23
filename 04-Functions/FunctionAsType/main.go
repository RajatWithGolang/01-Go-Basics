//A function in Go is also a type. 
//If two function accepts the same parameters and returns the same values,
// then these two functions are of the same type.
package main

import "fmt"

type myFunc func(int,int) int // it means myFunc is a type function which takes two arguments 
                               // and returns and int

func add(a,b int) int{
	return a+b
}

func sub(a,b int) int{
	return a-b
}

func calc(a,b int, f myFunc) int{
   r := f(a,b)
   return r
}
func main(){
 addresult := calc(10,5,add)
 subresult := calc(10,9,sub)
 fmt.Println(addresult)
 fmt.Println(subresult)
}

