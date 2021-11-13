package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hey there"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our movie: ")

	//comma ok syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating!,", input)

}
