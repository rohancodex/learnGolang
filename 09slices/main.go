package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices!")

	var fruitlist = []string{"apple", "mango", "papaya"}
	fmt.Println(fruitlist)

	//add data
	fruitlist = append(fruitlist, "Banana")
	fmt.Println(fruitlist)

	//slice your slices
	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScore := make([]int, 5)

	highScore[0] = 255
	highScore[1] = 255
	highScore[2] = 255
	highScore[3] = 255
	highScore[4] = 255
	// highScore[5] = 233

	//reallocation
	highScore = append(highScore, 233, 2323)
	fmt.Println(highScore[5])
	sort.Ints(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println(highScore)

	//remove values from slices
	index := 2
	var languages = []string{"javascript", "golang", "java", "python"}
	fmt.Println(languages)
	languages = append(languages[:index], languages[index+1:]...)
	fmt.Println(languages)
}
