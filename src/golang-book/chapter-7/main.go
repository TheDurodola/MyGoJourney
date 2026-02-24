package main

import "fmt"

func main() {


	fmt.Printf("add(1, 2): %v\n", add(1, 2))
	fmt.Printf("subtract(45.34, 6): %v\n", subtract(45.34, 6))

	

	multiplyResult, divideResult := multiply(10, 5)
	fmt.Printf("multiply(10, 5): %v\n", multiplyResult)
	fmt.Printf("divide(10, 5): %v\n", divideResult)


	fmt.Printf("min(10, 20, 30, 40): %v\n", min(10, 20, 30, 40))
	


	 max := func(args ...int) int {
		 maxValue := args[0]
		 for _, value := range args {
			 if value > maxValue {
				 maxValue = value
			 }
		 }
		 return maxValue
	 }

	 fmt.Printf("max(10, 20, 30, 40): %v\n", max(10, 20, 30, 40))

	 var number1 int8 = 1

	 increment := func() int8{
		 number1++
		 return number1
	 }

	 defer fmt.Printf("deferred number: %v\n", number1)


	 fmt.Printf("increment(): %v\n", increment())
	 fmt.Printf("increment(): %v\n", increment())

}


func add(x, y int) int {
	return x + y
}

func subtract(x float64, y int) float64 {
	return x - float64(y)
}

func multiply(x, y int) (float64, float64){
	var newFirstNumber float64 = float64(x)
	var newSecondNumber float64 = float64(y)

	multiplyResult := newFirstNumber * newSecondNumber
	divideResult := newFirstNumber / newSecondNumber

	return multiplyResult, divideResult
} 


func min(args ...int) int {
	if len(args) == 0 {
		panic("min function requires at least one argument")
	}

	minValue := args[0]
	for _, value := range args {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}