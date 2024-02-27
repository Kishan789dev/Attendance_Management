package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kk/attendance_management/dataBase"
	"github.com/kk/attendance_management/routers"

	"github.com/rs/cors"
)

type App struct {
	r *routers.RouteImp
	// db       *pg.DB
	database dataBase.DataBase
}

func NewApp(r *routers.RouteImp, database dataBase.DataBase) *App {
	return &App{
		r: r,
		// db:       db,
		database: database,
	}
}

func (app *App) Start() {

	r := mux.NewRouter()
	// _ = app.database.Connect()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	app.r.InitialiseRouter(r)
	defer log.Fatal(http.ListenAndServe(":9800", handler))
	defer app.database.Close()
}
