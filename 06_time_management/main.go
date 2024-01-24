package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time management with golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) 
	// * This is how we format time in go dd-mm-yyyy or any format with the given date, same with time, and day should be monday
	createdDate := time.Date(2023, time.May,1,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
}