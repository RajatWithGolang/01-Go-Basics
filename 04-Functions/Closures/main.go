package main

import "fmt"

type add func(a int,b int) int

func main(){
 var a add = func(a int, b int) int {
        return a + b
    }
 fmt.Println(a(10,5))
}