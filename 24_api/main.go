package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int  `json:"coursePrice"`
	Author      *Author `json:"author"`
}

// Model for Author - file
type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// helper - file
func (c *Course) IsEmpty() bool {
	return  c.CourseName == ""
}

// fake DB
var courses []Course

func main() {
	r := mux.NewRouter()
	courses = append(courses,Course{CourseId: "2",CourseName: "New1", CoursePrice: 200, Author: &Author{Fullname: "Anuj", Website: "https://github.com/anuj2110"}})
	fmt.Println(courses)
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/courses",createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/courses/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")
	http.ListenAndServe(":4000",r)
}

// controllers - files

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by Anuj</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Sending all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
	return 
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get course by Id")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,course := range(courses) {
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return 
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Creating the course")
	w.Header().Set("Content-Type","application/json")
	if r.Body==nil{
		json.NewEncoder(w).Encode("Empty Body")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("courseName is empty")
	}
	rand.Seed(time.Now().UnixNano())
	randomInt := strconv.Itoa(rand.Intn(100))
	course.CourseId = randomInt
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return 
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update the course")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,course := range(courses){
		if course.CourseId==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return 
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete the course")
	params := mux.Vars(r)
	for index,course := range(courses){
		if course.CourseId==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return 
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return 
}