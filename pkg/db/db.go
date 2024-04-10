package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id int
	Price float64
}


func LoadItem(id int) *Item {
	return &Item{
		Id: id,
		Price: 9.001,
	}
}

func DbConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./testDatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
	return db
}