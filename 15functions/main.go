package main

import (
	"fmt"
)

func main() {
	fmt.Println("functions in golang")
	ans := adder(5, 3)
	answer := proadder(10, 22, 16, 45, 879, 44)
	fmt.Println(ans)
	fmt.Println(answer)
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

//multiple values in parameter all of type int
func proadder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
