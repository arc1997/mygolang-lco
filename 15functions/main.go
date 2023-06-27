package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greet("Adesh")

	result := Add(3, 6)

	fmt.Println(result)

	fmt.Println(sliceAdd(1, 2, 4))
	one, two, three := multiple()
	fmt.Println(one, two, three)
}

func greet(name string) {
	fmt.Println("Namaste " + name + " in golang")
}

func Add(val1 int, val2 int) int {
	return val1 + val2
}

func sliceAdd(values ...int) int {
	sum := 0

	for _, val := range values {
		sum += val
	}
	return sum
}

// function return multiple values

func multiple() (int, string, int) {
	return 0, "from multiple", 0
}
