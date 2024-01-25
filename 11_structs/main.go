package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// * No inheritance in golang
	anuj := User{"Anuj","atrehan789@gmail.com",16,true}
	fmt.Println("This is my Struct -> ",anuj)
	fmt.Printf("These are the values with name %+v\n",anuj)
	fmt.Println(anuj.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
