package dataBase

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

type DataBase interface {
	Connect() *pg.DB
	Close() error
}

type DataBaseImpl struct {
	db *pg.DB
}

func NewDataBase() *DataBaseImpl {
	//

	// opts := &pg.Options{
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Addr:     os.Getenv("DB_ADD"),
	// 	Database: os.Getenv("DB_DATABASE"),
	// }
	// fmt.Println("test", opts)
	// fmt.Println("test", opts.User)
	db := ConnectTest()
	return &DataBaseImpl{
		db: db,
	}
}

func ConnectTest() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	fmt.Println("test2", opts)
	db := pg.Connect(opts)

	if db == nil {
		// log.Printf("database connection failed.\n")
		os.Exit(100)
	}
	// log.Printf("connect successful.\n")
	// cnt = db
	return db
}

func (impl *DataBaseImpl) Connect() *pg.DB {
	// if cnt != nil {
	// 	return cnt
	// }
	// opts := &pg.Options{
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Addr:     os.Getenv("DB_ADD"),
	// 	Database: os.Getenv("DB_DATABASE"),
	// }
	// fmt.Println("test2", opts)
	// db := pg.Connect(opts)

	// if db == nil {
	// 	// log.Printf("database connection failed.\n")
	// 	os.Exit(100)
	// }
	// // log.Printf("connect successful.\n")
	// // cnt = db
	// return db
	// fmt.Print(impl.db)
	return impl.db
}
func (impl *DataBaseImpl) Close() error {
	// if cnt != nil {
	// 	return cnt
	// }
	// opts := &pg.Options{
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Addr:     os.Getenv("DB_ADD"),
	// 	Database: os.Getenv("DB_DATABASE"),
	// }

	// db := pg.Connect(opts)

	// if db == nil {
	// 	log.Printf("database connection failed.\n")
	// 	os.Exit(100)
	// }
	// log.Printf("connect successful.\n")
	// cnt = db
	// return db

	return impl.db.Close()
}
