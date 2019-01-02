package main

import "fmt"

func main() {
	// Define maps
	//emails := make(map[string]string)

	// Assign key values
	// emails["Satya"] = "satya@gmail.com"
	// emails["Bob"] = "bob@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	//Declare and add KV

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//delete one from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
