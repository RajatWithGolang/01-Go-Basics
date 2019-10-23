package main

import "fmt"

func main(){
  arr := [9]int{1,2,3,4,5,6,7,8,9}
  fmt.Println("items in an array are :",arr)
  xi := arr[2:4]
  fmt.Println("items in a slice are :",xi)

 // Output :
// items in an array are : [1 2 3 4 5 6 7 8 9]
//items in a slice are : [3 4]

//==============================================

//let's modify an array
  arr[2] =33
  fmt.Println("Array After chaned :",arr)
  fmt.Println("Length of an arry is :",len(arr)) 
  fmt.Println("items in a slice also changes to :",xi)
   fmt.Println("length of a slice are :",len(xi))
   fmt.Println("capacity of a slice is :",cap(xi))
 // Output :
// items in an array are : [1 2 3 4 5 6 7 8 9]
// items in a slice are : [3 4]
// items in changed array are : [1 2 33 4 5 6 7 8 9] 
// Length of an arry is : 9
// items in a slice are : [33 4]  //changes is done in slice as well
// length of a slice are : 2
// capacity of a slice is : 7 --> start from arry index position 2 to 9 

//================================================

//let's modify the Slice
  xi[1] = 44
   fmt.Println("Slice is changed to :",xi)
  fmt.Println("Array after Slice change :",arr) // this also got change
  fmt.Println("Length of an arry is :",len(arr)) 

//Output:
// Slice is changed to : [33 44]
// Array after Slice change : [1 2 33 44 5 6 7 8 9]  --> array also got change
// Length of an arry is : 9

 
}	