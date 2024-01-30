package main

import (

	"encoding/json"
	"fmt"
)
type person struct{
	Name string
	Age int
	Address string
}
func main() {
	fmt.Println("Working with json")
	makeJson()
	readJson()
}

func makeJson(){
	data := []person{
		{"Anuj",25,"Vikas Puri"},
		{"Akash",25,"Noida"}}

	jsonData,_ := json.MarshalIndent(data,"","\t")
	fmt.Println(string(jsonData))
}

func readJson(){
	data := []byte(
		` {
			"Name": "Anuj",
			"Age": 25,
			"Address": "Vikas Puri"
	}`)

	jsonIsValid := json.Valid(data)
	
	if jsonIsValid{
		var myjson person
		json.Unmarshal(data,&myjson)
		fmt.Println("here:",myjson)
		var myUnknownJson map[string]interface{}
		json.Unmarshal(data,&myUnknownJson)
		fmt.Println("here:",myUnknownJson)
		for k,v := range(myUnknownJson){
			fmt.Printf("key is %v, value is %v and type is %T\n",k,v,v)
		}
	}else{
		fmt.Println("Data is not valid")
	}
	
}