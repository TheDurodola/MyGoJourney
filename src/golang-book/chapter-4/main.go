package main

import("fmt")

const pi float64 = 3.14159

func main(){
	var x uint8 = 255
	var v string = "Hello, World!"

	

	fmt.Printf("The value of pi is: %f\n", pi)
	fmt.Println(v)
	fmt.Println(x)

	var name string = "Bolaji"
	v += name
	fmt.Println(v)
	println("This is a println function")


	addedPi := addToPi(2.71828)
	fmt.Printf("The value of pi + e is: %f\n", addedPi)


	var (
		firstNumber int = 10
		secondNumber int = 20
		sum int = firstNumber + secondNumber
	)
	fmt.Printf("The sum of %d and %d is %d\n", firstNumber, secondNumber, sum)


	const (
		piValue float64 = 3.14159
		eValue float64 = 2.71828
	)
	fmt.Printf("The value of pi is: %f and the value of e is: %f\n", piValue, eValue)

	printUserNumber()
}

func println(s string){
	fmt.Println(s)
}


func addToPi(x float64) float64 {
	return pi + x
}

func printUserNumber(){
fmt.Println("Please enter a number:")
var userInput int
fmt.Scanln(&userInput)
fmt.Printf("You entered: %d\n", userInput)
}
