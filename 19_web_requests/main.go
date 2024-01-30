package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"
func main() {
	fmt.Println("Web Requests")
	response,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("This is the web resonse %T\n",response)
	defer response.Body.Close()
	dataBytes,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content:=string(dataBytes)
	fmt.Println(content)
}