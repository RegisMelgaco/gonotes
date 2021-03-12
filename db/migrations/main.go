package main

import (
	"fmt"

	"github.com/go-pg/migrations"

	"local/gonotes/db"
)

func main() {
	dbConnection := db.GetDbConnection()
	defer dbConnection.Close()

	migrations.Run(dbConnection, "init")
	oldVersion, nextVersion, err := migrations.Run(dbConnection, "up")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Updated from ", oldVersion, " to ", nextVersion)
	}
}
