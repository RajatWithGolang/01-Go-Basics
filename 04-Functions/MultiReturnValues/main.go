package main 
import "fmt"
func details()(int,string){//Named return Values : func details()(age int,name string){
	return 29,"Rajat"
	//return -->Necesssary
}
func main(){
	age,name:= details()
	fmt.Println(age,name)
}