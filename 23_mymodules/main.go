package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Making first go server with mux")
	r := mux.NewRouter()
    r.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":4000",r))
}

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello from gorilla mux</h1>"))
}