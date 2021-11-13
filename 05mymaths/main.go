package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("All about Math in golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4.5

	//generates random number from 0 to 4

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	randomNo, _ := rand.Int(rand.Reader, big.NewInt(6))

	fmt.Println(randomNo)

}
