package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	database "github.com/kk/attendance_management/dataBase"
	"github.com/kk/attendance_management/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	defer log.Fatal(http.ListenAndServe(":9800", handler))
	// defer databaseconnection.Close()
	// log.Fatal(http.ListenAndServe(":9800", r))

}
