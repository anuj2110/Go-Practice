package main

import "fmt"

func main(){
	fmt.Println("Maps in Golang")
	languages:=make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("Languages are: ",languages)
	delete(languages,"RB")
	fmt.Println("Languages are: ",languages)
	for key,value:= range languages{
		fmt.Printf("For Key %v value is %v\n",key,value)
	}
}