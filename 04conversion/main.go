package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizza between 1 & 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numrating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to rating", numrating+1)
	}

}
