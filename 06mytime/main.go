package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time module of golang")

	presentime := time.Now().Month()
	fmt.Println(presentime)

	createdDate := time.Date(2020, time.April, 15, 20, 10, 10, 0, time.UTC)

	fmt.Println(createdDate)
}
