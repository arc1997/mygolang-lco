package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in golang")
	content := "This needs to be go in file by golang code"

	file, error := os.Create("./mygofile.txt")

	// if error != nil {
	// 	panic(error) // will shut the program & will throw the error
	// }
	CheckNillError(error)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	CheckNillError(err)

	fmt.Println("Length of file is : ", length)

	defer file.Close() // it will close after all operations done means at the end of the function
	ReadFile("./mygofile.txt")

}

func ReadFile(filename string) {
	databyte, error := os.ReadFile(filename)

	// if error != nil {
	// 	panic(error)
	// }
	CheckNillError(error)

	fmt.Println(databyte)

	fmt.Println(string(databyte))
	fmt.Printf("File contents: %s", databyte)
}

//Common function to panic the error

func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}
