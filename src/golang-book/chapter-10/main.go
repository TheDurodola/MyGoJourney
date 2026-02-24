package main

import (
	"fmt"
	"time"
	"math/rand"
)
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(250)
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {

	var x []string ;

	x = append(x, "rock")
	x = append(x, "paper")
	x = append(x, "scissors")



	for  i := 0; ;i++ {
		c <- x[rand.Intn(3)]
	} 
}

func printer(c chan string) {
	for{
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}



func main() {
	
	c := make(chan string)

	go printer(c)
	go pinger(c)


	var input string
	fmt.Scanln(&input)
}