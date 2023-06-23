package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	var one int = 1

	fmt.Println(one)

	//var ptr *int

	//fmt.Println(ptr)

	mynumber := 23

	var ptr = &mynumber
	// prints memory address
	fmt.Println(ptr)
	//prints value in that memory address
	fmt.Println(*ptr)

	*ptr = *ptr * 2

	fmt.Println(mynumber)
}
