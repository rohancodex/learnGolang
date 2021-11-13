package main

import "fmt"

func main() {
	fmt.Println("loops in golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	value := 1
	for value < 5 {

		if value == 4 {
			goto fun
		}

		if value == 3 {
			value++
			continue
		}
		fmt.Println("Value is: ", value)
		value++
	}

fun:
	fmt.Println("Fun Stuff")
}
