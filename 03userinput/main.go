package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name")

	input, _ := reader.ReadString('\n')

	fmt.Println("Entered name", input)
	fmt.Printf("Type of name is %T", input)
}
