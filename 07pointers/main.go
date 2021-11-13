package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Value of pointer is", ptr)

	myNumber := 23 //var myNumber = 23
	ptr = &myNumber
	fmt.Println("The value of pointer is ", ptr)
	fmt.Println("The value of pointer is ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value is: ", myNumber)
}
