package main

import "fmt"

func change(arr [9]int){
	ind := 2
	for index,val := range arr{
		arr[index] = val+ind
	}
	fmt.Println("array changed to ",arr)
}

func main(){
	arr:= [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println("array before change",arr)
	change(arr)
     fmt.Println("array after change",arr)

}

//When you pass an array to a function, they are passed by value like int or string data type.
// The function receives only a copy of it. 
//Hence, when you make changes to an array inside a function, it won’t be reflected in the original array.

//Output
// array before change [1 2 3 4 5 6 7 8 9]
// array changed to  [3 4 5 6 7 8 9 10 11]
// array after change [1 2 3 4 5 6 7 8 9]
