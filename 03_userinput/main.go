package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Rating for our Course: ")

	//comma ok syntax || err ok
	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("Type of the rating is %T",input)
}