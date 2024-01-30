package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("Hello")
	myDefer()
	/*
	* Hello
	* 2
	* 1
	* World
	*/
}

func myDefer(){
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
}