package db

import "github.com/go-pg/pg"

//GetDbConnection returns a connection with the aplication's DB
func GetDbConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
		Password: "mysecretpassword",
		Addr:     "192.168.1.88:5432",
	})
}
