package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	//var mynumberone int = 2
	//var mynumbertwo float64 = 4

	//fmt.Println("The sum is : ",mynumberone+int(mynumbertwo))

	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int31n(5))

	//random  number by crypto
	myrandom, _ := rand.Int(rand.Reader, big.NewInt(6))

	fmt.Println(myrandom)

}
