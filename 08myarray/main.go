package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Jamun"
	fmt.Printf("Type of fruitlist is %T \n", fruits)
	fmt.Println(fruits[2])
	fmt.Println(len(fruits))

	//initalize
	var integers = [3]int{1, 2, 3}

	fmt.Println(integers)

}
