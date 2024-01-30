package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// func initialiseRouter() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/student/{id}", getStudent).Methods("GET")
// 	r.HandleFunc("/student", getStudents).Methods("GET")
// 	r.HandleFunc("/student", addStudent).Methods("POST")
// 	r.HandleFunc("/student{id}", updateStudent).Methods("PUT")
// 	r.HandleFunc("/student{id}", deleteStudent).Methods("DELETE")

// 	// ************Teacher************************************

// 	r.HandleFunc("/techer/{id}", getTeacher).Methods("GET")
// 	r.HandleFunc("/teacher", getTeachers).Methods("GET")
// 	r.HandleFunc("/teacher", addTeacher).Methods("POST")
// 	r.HandleFunc("/teacher{id}", updateTeacher).Methods("PUT")
// 	r.HandleFunc("/teacher{id}", deleteTeacher).Methods("DELETE")

// 	// ********************Teacher attendance*****************

// 	r.HandleFunc("/teacherattendance", addTeacherattendace).Methods("PUT")
// 	r.HandleFunc("/teacherattendace/{id}/{month}/{year}", getTeacherattendace).Methods("GET")
// 	// principle
// 	// r.HandleFunc("/teacherattendace/{id}/{month}/{year}").Methods("GET")

// 	// ********************Student attendance*****************

// 	r.HandleFunc("/studentattendance", addStudentattendance).Methods("PUT")
// 	r.HandleFunc("/studentattendance/{id}/{month}/{year}", getStudentattendance).Methods("GET")
// 	r.HandleFunc("/studentattendance/{class}/{date}/{month}/{year}", getClassattendance).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":9000", r))
// }

func main() {
	// initialiseRouter()

	// product := Product{ID: uuid.New(), Name: "kishan", Quantity: 34}

	// fmt.Println(product)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}
	fmt.Println(uuid.New())
	db := connect()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/products", createProducts).Methods("POST")

	log.Fatal(http.ListenAndServe(":8005", r))

}

func createProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// get connect
	db := connect()
	defer db.Close()

	// creating product instance
	product := &Product{
		ID: uuid.New(),
	}

	// decoding request

	_ = json.NewDecoder(r.Body).Decode(&product)

	// inserting into database
	_, err := db.Model(product).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// returning product

	json.NewEncoder(w).Encode(product)

}
