package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // alias name in json data
	Price    int
	Platform string
	Password string   `json:"-"`              // - represents here that we do not want this field in json
	Tags     []string `json:"tags,omitempty"` // omitempty can be used to remove null values
}

func main() {
	fmt.Println("Welcome to json handling in go")
	Encodejson()
}

func Encodejson() {
	courses := []course{
		{"React js", 299, "LearnCode", "abc123", []string{"web-dev", "js"}},
		{"mern", 499, "LearnCode", "bcd123", []string{"full-stack", "js"}},
		{"mean", 799, "LearnCode", "abd123", nil},
	}

	// encode this data as json data
	//jsondata, err := json.Marshal(courses)
	jsondata, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsondata))
}
