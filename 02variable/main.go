package main

import "fmt"

const UserToken = "hkjjhk"

func main() {
	var username string = "Adesh"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isuser bool = true
	fmt.Println(isuser)
	fmt.Printf("variable is of type %T \n", isuser)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type %T \n", smallval)

	var smallfloat float64 = 255.0166464646
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type %T \n", smallfloat)

	// default values & aliases
	var variable bool
	fmt.Println(variable)
	fmt.Printf("variable is of type %T \n", variable)

	// implicit
	var variable1 = "adesh"
	fmt.Println(variable1)
	fmt.Printf("variable is of type %T \n", variable1)

	// no var style
	numberofuser := 3000
	fmt.Println(numberofuser)
}
