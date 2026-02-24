package main


import "fmt"


func main(){
	var x int8 = 0
	addOne(x)
	fmt.Println(x)
	addOneUsingPointer(&x)
	fmt.Println(x)
}

func addOne(x int8) int8 {
return x + 1
}

func addOneUsingPointer(x *int8){

	*x = 8

}