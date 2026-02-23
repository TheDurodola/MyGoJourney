package main

import "fmt"

func main(){

	// printArray()
playingWithSlices()
playingWithMaps()

}


func printArray(){
	var array [5]int
	array[0] = 10
	fmt.Println(array)
}


func playingWithSlices(){
	var slice []uint16;

	fmt.Println(slice)

	slice = append(slice, 3)
	fmt.Println(slice)
	slice = append(slice, 4, 5)
	fmt.Println(slice)
	fmt.Printf("The type is: %T\n", slice)
	copy(slice, []uint16{1, 2, 3})
	fmt.Println(slice)
} 

func playingWithMaps(){
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	fmt.Println(map1)

	delete(map1, "one")
	fmt.Println(map1)
	fmt.Println(map1["HMM"])
	fmt.Println(map1["HMM"]==0)

	name, ok := map1["Un"]
fmt.Println(name, ok)
}