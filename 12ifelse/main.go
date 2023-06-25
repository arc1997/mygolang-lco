package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang")

	logincount := 10

	var result string
	if logincount < 10 {
		result = "Regular user"
	} else if logincount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("More than 10")
	}

	// if err != nil {

	// }
}
