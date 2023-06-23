package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println(languages["JS"])

	delete(languages, "RB")

	fmt.Println(languages)

	//loops are interesting in golang

	for key, value := range languages {
		//fmt.Printf("For key %v, the value is: %v\n", key, value)
		fmt.Println(key + ": " + value)
	}
}
