package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	rohan := User{"Rohan", "rohan@email.com", true, 20}

	fmt.Println(rohan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
