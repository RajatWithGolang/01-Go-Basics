package main

import "fmt"

func change(xi []int){
	ind := 2
	for index,val := range xi{
		xi[index] = val+ind
	}
	fmt.Println("array changed to ",xi)
}

func main(){
	xi:= []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Slice before change",xi)
	change(xi)
     fmt.Println("Slice after change",xi)
//Output :
// Slice before change [1 2 3 4 5 6 7 8 9]
// array changed to  [3 4 5 6 7 8 9 10 11]
// Slice after change [3 4 5 6 7 8 9 10 11]

}

// slices are still passed by value to the function but since they reference the same underneath array
// it looks like that they are passed by reference.

