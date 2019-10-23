package main

import "fmt"

func findNumber(num int,nums ...int){
	 found :=false
	 for ind,val := range nums{
		 if val == num{
			 fmt.Printf("%d is found at %d index",num,ind)
			 found = true
		 }
		}
		if !found{
		 fmt.Println("% does not exist")
		}	
}

func main(){
      findNumber(10,20,30,40,50,10,70,80)
}