package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
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
	//---------------- without go
	start := time.Now()
	duration := time.Since(start)
	for i := 0; i < 4; i++ {
		inserterSimple(i + 20)
	}
	fmt.Println("without go - " + fmt.Sprintf("%.6f", duration))

	//---------------- with go
	db, err := sql.Open("postgres", getDsn())
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()

		if err != nil {
			panic(err)
		}
	}()

	var wg sync.WaitGroup
	startGo := time.Now()
	durationGo := time.Since(startGo)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go inserter(i+20, &wg, db)
	}
	wg.Wait()
	fmt.Println("with go - " + fmt.Sprintf("%.6f", durationGo))

}

func getDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func getConnect() *sql.DB {
	db, err := sql.Open("postgres", getDsn())
	db.SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}
	return db
}

func inserter(num int, group *sync.WaitGroup, db *sql.DB) {
	defer group.Done()

	log.Print(db.Stats().OpenConnections)
	sqlStatement := `
		INSERT INTO data (name, description)
		VALUES ($1, $2)`

	for i := 0; i < 1000; i++ {
		_, err := db.Exec(sqlStatement, num, num)
		if err != nil {
			panic(err)
		}
	}
}

func inserterSimple(num int) {
	db := getConnect()
	defer func() {
		err := db.Close()

		if err != nil {
			panic(err)
		}
	}()

	sqlStatement := `
		INSERT INTO data (name, description)
		VALUES ($1, $2)`

	for i := 0; i < 1000; i++ {
		_, err := db.Exec(sqlStatement, num, num)
		if err != nil {
			panic(err)
		}
	}
}
