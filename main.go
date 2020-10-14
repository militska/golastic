package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
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

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go inserter(i+20, &wg)
	}

	wg.Wait()
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

func inserter(num int, group *sync.WaitGroup) {
	db := getConnect()
	defer func() {
		group.Done()
		err := db.Close()

		if err != nil {
			panic(err)
		}
	}()

	sqlStatement := `
		INSERT INTO data (name, description)
		VALUES ($1, $2)`

	for i := 0; i < 4000; i++ {
		_, err := db.Exec(sqlStatement, num, num)
		if err != nil {
			panic(err)
		}
	}
}
