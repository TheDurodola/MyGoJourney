package main

import "fmt"

func main() {
	controlStructure()


}


func controlStructure() {
	forLoop()
	ifElse()
	switchStatement()
	
}

func forLoop() {
const count int = 10

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}


func ifElse() {
const age int = 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
	
}


func switchStatement() {
const day int = 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}
