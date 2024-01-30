package main

import (

	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content:="This content will go in file"
	file,err := os.Create("./my_txt_file.txt")
	
	if err!=nil{
		panic(err)
	}
	length,err := io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Length is %v",length)
	defer file.Close()
	readFile("./my_txt_file.txt")
}

func readFile(fileName string){
	dataByte,err := ioutil.ReadFile(fileName)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The data is: ",string(dataByte))

}