package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(3)+1

	switch diceNumber {
	case 1:
		fmt.Println("You can open")
		fallthrough
	case 2:
		fmt.Println("2 Steps")
	case 3:
		fmt.Println("3 steps")
	default:
		fmt.Println("What was that")
	}
}