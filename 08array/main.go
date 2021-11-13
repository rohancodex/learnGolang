package main

import "fmt"

func main() {
	var fruitlist [4]string
	fmt.Println("arrays!")
	fruitlist[0] = "Apple"
	fruitlist[1] = "Guava"
	fruitlist[2] = "Papaya"

	fmt.Println("Fruit list is: ", fruitlist)

	var veggies = [4]string{"beans", "spinach", "carrot"}
	fmt.Println("Veggies are: ", veggies)
}
