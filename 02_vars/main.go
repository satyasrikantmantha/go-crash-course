package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint 32 uint64 uintotr
	// byte - alias fo ruint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
	//var name = "Satya"
	//Shorthand variable syntax, only inside a function
	// name := "Satya"
	// email := "satya@gmail.com"
	var age int32 = 36
	const isCool = true
	var size float32 = 2.3

	//Multiple assignments with shorthand syntax
	name, email := "Satya", "satya@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
