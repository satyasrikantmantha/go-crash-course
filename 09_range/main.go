package main

import "fmt"

func main() {
	ids := []int{33, 54, 34, 67, 23, 86}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d -- ID : %d \n", i, id)
	}

	// Not using index, if you don't want to use index
	for _, id := range ids {
		fmt.Printf("ID : %d \n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum of Ids:", sum)

	//Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Access only Keys
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
