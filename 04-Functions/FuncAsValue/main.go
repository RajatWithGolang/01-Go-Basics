//A function in go can also be a value. This means you can assign a function to a variable

package main

import "fmt"

var add = func(a,b int) int{
	return a+b
}
func main(){
    fmt.Println(add(10,20))// add is an anonymous function
}

//immedialty invoked function
// func main() {
// 	sum := func(a int, b int) int {
// 		return a + b
// 	}(3, 5)
	
// 	fmt.Println("5+3 =", sum)
// }