package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url string = "https://jsonplaceholder.typicode.com/posts"
func main() {
	fmt.Println("Web requests")
	// makeGetRequest()
	makePostRequest()
	makeGetRequest()
}

func makeGetRequest(){
	response,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	data,err := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

}

func makePostRequest(){
	/*
	data:=url.Values{}
	data.Add("","")
	http.PostForm()
	*/
	data := strings.NewReader(`
		{
			"title": "foo",
    		"body": "bar",
    		"userId": 1
		}
	`)
	response,err := http.Post(url,"application/json",data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	bytes_,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes_))
}
