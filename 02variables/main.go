package main

import "fmt"

const LoginToken = "asjkdhkasd"

func main() {
	var username string = "Rohan"
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.23131
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases
	var someVariable bool
	fmt.Println(someVariable)

	//implicit way
	var website = "rohandesai.me"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
