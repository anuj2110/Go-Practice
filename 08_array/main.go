package main

import "fmt"

func main() {
	var myFruits [4]string;
	myFruits[0] = "Banana"
	myFruits[1] = "Kiwi"
	myFruits[3] = "Apple"
	fmt.Println("This is my First Array",myFruits)
	var myVeggies = [3]string{"Capsicum","Beans","Gobhi"}
	fmt.Println("This is my Second Array",myVeggies)
}