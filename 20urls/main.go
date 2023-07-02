package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=jgjdf8998"

func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println(myurl)

	//parsing the url using url parser
	result, error := url.Parse(myurl)

	if error != nil {
		panic(error)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of qparams is : %T\n", qparams)

	fmt.Println(qparams["coursename"][0])

	fmt.Println(result.Query().Get("paymentid"))

	for key, val := range qparams {
		fmt.Println("For key ", key, "value is : ", val)
	}

	//contruct URL

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "?user=hitesh",
	}

	anotherurl := partsofurl.String()

	fmt.Println(anotherurl)

}
