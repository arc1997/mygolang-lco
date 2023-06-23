package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitlist = []string{}

	fmt.Printf("Type of fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Apple", "Tomato")

	fruitlist = append(fruitlist, "Banana")
	fmt.Println(fruitlist)

	//fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	scores := make([]int, 4)

	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	scores[3] = 4
	//scores[4] = 5

	scores = append(scores, 6, 5)

	fmt.Println(scores)

	sort.Ints(scores)

	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	//how to remove value a value from slices based on  index

	var courses = []string{"react.js", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
