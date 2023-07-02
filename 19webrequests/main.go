package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/robots.txt"

func main() {
	fmt.Println("Web Requests in Golang")

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	//fmt.Println(response)

	fmt.Printf("Type of Response is : %T \n", response) //type => *http.Response

	defer response.Body.Close() // caller's responsibility to close the connection.

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))

}
