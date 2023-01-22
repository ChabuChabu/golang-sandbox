package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apples"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list is", len(fruitList))

	var vegList = [3]string{"potatoes", "Beans", "mushroom"}

	fmt.Println("Veg list is:", vegList)
}
