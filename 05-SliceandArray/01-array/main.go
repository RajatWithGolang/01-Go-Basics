//An array is a container that holds data of same type. 
//Arrays in Go have fixed length and they can’t be expanded to fit more data

package main

import "fmt"

func main(){
	//===========Array Declaration==============
 //var arr[3] int
// arr := [3]int{} -> shorthand syntax
 //fmt.Println(arr)// it will have zero value of int data type
 //============Declaration=======================
 arr:= [3]string{
	 "Rajat",
	 "Rishabh",
	 "Richa",// must have comma
 }
 for ind,value := range arr{
	 fmt.Printf("%s is present at index %d\n",value,ind)
}
}

