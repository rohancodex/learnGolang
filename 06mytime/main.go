package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("all about time")

	//formatting the present time
	fmt.Println(time.Now().Format("01-02-2006 15:04:05"))

	t := time.Now()

	t.Year()

	t.Month()

	t.Day()

	t.Hour()

	t.Minute()

	t.Second()

}
