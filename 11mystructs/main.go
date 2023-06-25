package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance
	// No super or parent

	user := User{"Adesh", "askadeshchhajed@gmail.com", true, 26}

	fmt.Println(user)

	fmt.Printf("user details is : %+v \n", user)

	fmt.Println(user.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
