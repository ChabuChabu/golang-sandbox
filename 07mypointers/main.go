package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers!!")

	// var pointer *int
	// fmt.Println("The of pointer is:", pointer)

	myNumber := 23

	var pointer = &myNumber

	fmt.Println("Value of actual pointer is", pointer)
	fmt.Println("Value of actual pointer is", *pointer)

	*pointer = *pointer * 2

	fmt.Println("New value is", myNumber)
}
