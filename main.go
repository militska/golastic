package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "postgresql"
	port     = 5432
	user     = "user"
	password = "123"
	dbname   = "db_main"
)

func main() {
	v1() // без горутин
	v2() // с горутинами в1
	//v3() //с горутинами  в2

}

func getDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func getConnect() *sql.DB {
	db, err := sql.Open("postgres", getDsn())

	if err != nil {
		panic(err)
	}

	return db
}
