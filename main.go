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

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}
	fmt.Print("hello")
	db := database.Connect()
	defer db.Close()

	r := mux.NewRouter()
	routers.InitialiseRouter(r)
	log.Fatal(http.ListenAndServe(":9800", r))

}
