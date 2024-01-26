package main

import "fmt"

func main() {
	fmt.Println("Lets go With functions")
	greeter()
	result := adder(123,31231)
	fmt.Println(result)
	fmt.Println(proAdder(23,324,5234,62352))
	var thisArray [4]string
	fmt.Printf("Type of thisArray %T\n",thisArray)
	thisSlice:= []string{}
	fmt.Printf("Type of thisMap %T",thisSlice)
}

func adder(val1 int,val2 int) int {
	return val1+val2
}

func proAdder(values ...int) int {
	total:=0
	for _,val := range values{
		total+=val
	}
	return total
}
func greeter(){
	fmt.Println("Hello from greeter")
}