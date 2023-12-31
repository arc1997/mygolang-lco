package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests verb")
	//GetRequest()
	//PostRequest()
	FormDataRequest()
}

func GetRequest() {
	const myurl = "http://localhost:8000/get"

	response, error := http.Get(myurl)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	var responseString strings.Builder
	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	bytecount, _ := responseString.Write(databyte)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())

}

func PostRequest() {
	const myurl = "http://localhost:8000/post"

	// dummy json load

	requestBody := strings.NewReader(`
	{
		"coursename" : "golang",
		"paymentid" : "jhkd897",
		"price" : 500
	}
	`)

	response, error := http.Post(myurl, "application/json", requestBody)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	var responseString strings.Builder
	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	bytecount, _ := responseString.Write(databyte)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())

}

func FormDataRequest() {
	const myurl = "http://localhost:8000/postform"

	//dummy form data

	formdata := url.Values{}
	formdata.Add("course", "golang")
	formdata.Add("price", "256")

	response, error := http.PostForm(myurl, formdata)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	var responseString strings.Builder
	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	bytecount, _ := responseString.Write(databyte)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())

}
