package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])

	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println("index is : " + strconv.Itoa(index) + " value is : " + day)
	// }

	i := 1
	// break statement

	// for i < 10 {
	// 	if i == 5 {
	// 		i++
	// 		break

	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// continue statement
	// for i < 10 {
	// 	if i == 5 {
	// 		i++
	// 		continue

	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// goto statement

	i = 1
	for i < 10 {
		if i == 5 {
			i++
			goto lco

		}
		fmt.Println(i)
		i++
	}
lco:
	fmt.Println("Jumping here")
}
