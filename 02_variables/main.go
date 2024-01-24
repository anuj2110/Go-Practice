package main

import "fmt"

// jwtToken := 4324234234 -> Invalid outside function bodies
const LoginToken string = "sifoewjirjfewpf23e2" // -> Capital letter denotes public visibility
func main()  {
	var username string = "Anuj"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n",smallValue)

	var smallFloatValue float32 = 255.346827890237423
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n",smallFloatValue)

	// no type
	var something = "something"
	fmt.Println(something)
	fmt.Printf("Variable is of type: %T \n",something)
	// something = 2 -> error


	// no value
	var someone int //gives default value of the type
	fmt.Println(someone)
	fmt.Printf("Variable is of type: %T \n",someone)

	//walrus operator
	noOfUsers := 234233242
	fmt.Println(noOfUsers)
	fmt.Printf("Variable is of type: %T \n",noOfUsers)
}