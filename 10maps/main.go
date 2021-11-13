package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("list of languages ", languages)
	fmt.Println("JS stands for", languages["JS"])

	//delete values
	delete(languages, "JS")
	fmt.Println("list of languages ", languages)

	//looping thr maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
