package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// type mainImpl struct {
// 	route routers.Route
// }

//	func NewMain(route routers.Route) *mainImpl {
//		return &mainImpl{route: route}
//	}

func main() {
	// fmt.println("lkdjfsdklfjsdklfj")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}
	fmt.Print("hello")
	app := InitialiseApp()
	app.Start()

}
