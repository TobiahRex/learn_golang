package main

import (
	"fmt"
	"github.com/TobiahRex/learn_golang/pkg/db"
)

func main() {
	fmt.Println("Connect to DB: ")
	db.DbConnect()
}