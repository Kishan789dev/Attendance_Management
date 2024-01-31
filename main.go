package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	database "github.com/kk/attendance_management/dataBase"
	"github.com/kk/attendance_management/routers"
)

func main() {

	// product := Product{ID: uuid.New(), Name: "kishan", Quantity: 34}

	// fmt.Println(product)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}
	fmt.Print("hello")
	db := database.Connect()
	defer db.Close()

	// fmt.Println(uuid.New())
	//
	r := mux.NewRouter()
	routers.InitialiseRouter(r)
	log.Fatal(http.ListenAndServe(":9800", r))
	// r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	// r.HandleFunc("/api/v1/products", createProducts).Methods("POST")
	// r.HandleFunc("in/api/v1/products/{id}", getProduct).Methods("GET")
	// r.HandleFunc("/api/v1/products/{id}", updateProduct).Methods("PUT")

	// r.HandleFunc("/api/v1/products/{id}", deleteProduct).Methods("DELETE")

}
