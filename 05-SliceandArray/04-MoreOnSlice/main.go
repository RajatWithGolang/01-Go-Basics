package main

import "fmt"

func main(){
  arr := [9]int{1,2,3,4,5,6,7,8,9}
  fmt.Println("items in an array are :",arr)
  slice := arr[2:4]
  fmt.Println("items in a slice are :",slice)

  //=====using append===========
    //newSlice := append(slice,55,66)
    //fmt.Println("items in new slice are :",newSlice)
	//fmt.Println("items in slice are :",slice)
	//fmt.Println("items in an array are :",arr)
// Output :
// items in new slice are : [3 4 55 66] 
// items in slice are : [3 4] // --> append does not mutate original slice
// items in an array are : [1 2 3 4 55 66 7 8 9] --> append function mutated array referenced by slice s

//==================Use append only to self assign the new slice like s = append(s, ...) 

   slice = append(slice,33,44,55,66,77)
   fmt.Println("slice after append",slice)
   fmt.Println("length of slice",len(slice))
   fmt.Println("capacity of slice",cap(slice))
// Output :
// slice after append [3 4 33 44 55 66 77]
// length of slice 7
// capacity of slice 7

//=============exceed the Slice Capacity=================
   slice = append(slice,88)
   fmt.Println("slice after append",slice)
   fmt.Println("New length of slice",len(slice))
   fmt.Println("New capacity of slice",cap(slice))

  //Output :
// slice after append [3 4 33 44 55 66 77 88]
// New length of slice 8
// New capacity of slice 14  // capacity gets doubled




}