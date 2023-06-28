package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	user := User{"Adesh", "askadeshchhajed@gmail.com", true, 26}

	fmt.Println(user)

	fmt.Printf("user details is : %+v \n", user)

	fmt.Println(user.Name)

	user.Getstatus()
	user.NewMail()
	fmt.Printf("user details is : %+v \n", user)
}

type User struct {
	Name   string // 1st capital letters are public
	Email  string
	Status bool
	Age    int
	//oneAge int // Not exposed because 1st letter is not capital

}

func (user User) Getstatus() {
	fmt.Println("Is user active: ", user.Status)
}

func (user *User) NewMail() {
	user.Email = "adesh@gmail.com"

	fmt.Println(user.Email)
}
