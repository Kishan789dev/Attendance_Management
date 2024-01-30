package dataBase

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	// rh "github.com/kk/attendance_management/restHandler"
)

// database connection

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("database connection failed.\n")
		os.Exit(100)
	}
	log.Printf("connect successful.\n")

	// if err := createSchema(db); err != nil {
	// 	log.Fatal(err)
	// }
	return db

}

// func createSchema(db *pg.DB) error {
// 	models := []interface{}{
// 		(*Student)(nil),
// 	}

// 	for _, model := range models {
// 		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			IfNotExists: true,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 	}
// 	return nil

// }
