package main

import (
	"fmt"
	"net/url"
)

const urlString string = "https://google.com/something?name=anuj&age=25"
func main() {
	fmt.Println("Working with URLs")
	worker,err := url.Parse(urlString)

	if err!=nil{
		panic(err)
	}

	fmt.Println(worker.Host)
	fmt.Println(worker.Scheme)
	fmt.Println(worker.Path)
	fmt.Println(worker.RawQuery)
	queryParams := worker.Query()
	for key,val := range(queryParams){
		fmt.Println(key,val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "somethign",
	}

	mainUrl := partsOfUrl.String()
	fmt.Println(mainUrl)
}