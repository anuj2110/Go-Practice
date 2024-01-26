package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Some learning on slices")
	var fruits = []string{"Kiwi", "Banana", "Apple"}
	/*
		fruits[0] = "Apple"
		fruits[1] = "Banana"
		fruits[3] = "Kiwi"
		fmt.Println("Fruits -> ",fruits)
		* this will give error because the slice is intialized with 0 length/elements
	*/
	fruits = append(fruits, "Mango", "Watermelon")
	fmt.Println("Fruits -> ", fruits)

	// * slicing (no Pun)
	fruits = append(fruits[1:])
	fmt.Println("Fruits -> ",fruits)
	fruits = append(fruits[1:5])
	fmt.Println("Fruits -> ",fruits)

	//* make -> pointer approach
	highScores := make([]int,4)
	highScores[0] = 342
	highScores[1] = 234
	highScores[2] = 998
	highScores[3] = 456
	highScores = append(highScores, 324,456,674)
	fmt.Println("HighScores -> ",highScores)
	sort.Ints(highScores)
	fmt.Println("highScores -> ",highScores)
	index := 2
	highScores = append(highScores[:index],highScores[index+1:]...) // * -> This is spread like operator
	fmt.Println("highScores -> ",highScores)
}
