package main

import "fmt"

func main() {
	// //Arrays
	// var fruitArr [2]string

	// //Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// //Declare and assign in one go
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)

	//length
	fmt.Println(len(fruitSlice))

	//range
	fmt.Println(fruitSlice[1:2])
	fmt.Println(fruitSlice[1:3])
}
