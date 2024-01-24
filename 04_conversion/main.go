package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza app")
	fmt.Println("Please provide rating on scale of 1-5:")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	rating,err := strconv.ParseFloat(strings.TrimSpace(input),64) // * -> If use input only then gives error as the last \n is getting appended
	if err!=nil{
		fmt.Println("=========>",err)
	}else{
		fmt.Println("Your Rating+1 is: ",rating+1)
	}
}