package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var courses = []Course{{"1", "Golang", 400, &Author{"Adesh Chhajed", "www.adesh.com"}},
	{"2", "React", 500, &Author{"Adesh Chhajed", "www.adesh.com"}},
}

//middleware ,helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""

}
func main() {
	r := mux.NewRouter() // New router
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{courseid}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{courseid}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{courseid}", DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//controllers - file

// serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in golang</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through the courses, find matching id

	for _, val := range courses {
		if val.CourseId == params["courseid"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in json")
		return
	}

	// generate unique id ,string
	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	//add new course to courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through the courses, find matching id and remove the course & add updated course
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	for index, val := range courses {
		if val.CourseId == params["courseid"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}
	course.CourseId = params["courseid"]

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through the courses, find matching id and remove the course

	for index, val := range courses {
		if val.CourseId == params["courseid"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(params["courseid"])

}
