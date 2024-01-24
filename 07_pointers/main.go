package main

import "fmt"

func main() {
	var myPointer *int;
	fmt.Println("This is my first pointer ",myPointer)
	myNumber := 32
	var newPointer *int = &myNumber
	fmt.Println("This is the new pointer",newPointer)
	fmt.Println("This is the value of new pointer",*newPointer)
	*newPointer = *newPointer+3
	fmt.Println("Value of myNumber: ",myNumber)
}