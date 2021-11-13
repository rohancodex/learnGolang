package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a rating !")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	//convert string from input variable and trim the blank space to float number
	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Thanks for the ratings!,", numRating+1)

}
