package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "postgresql"
	port     = 5432
	user     = "user"
	password = "123"
	dbname   = "db_main"
)

func main() {
	log.Print("hi!")
	db := getConnect()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	go inserter(22)
	go inserter(44)
	go inserter(55)
	go inserter(60)

	time.Sleep(60 * time.Second)
}

func getConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func inserter(num int) {
	db := getConnect()
	defer db.Close()

	sqlStatement := `
		INSERT INTO data (num)
		VALUES ($1)`

	for i := 0; i < 2000; i++ {
		_, err := db.Exec(sqlStatement, num)
		if err != nil {
			panic(err)
		}
	}
}
