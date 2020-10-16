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
	////---------------- without go
	//start := time.Now()
	//duration := time.Since(start)
	//for i := 0; i < 4; i++ {
	//	inserterSimple(i + 20)
	//}
	//fmt.Println("without go - " + fmt.Sprintf("%.6f", duration))
	//
	////---------------- with go
	//db, err := sql.Open("postgres", getDsn())
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	err := db.Close()
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//var wg sync.WaitGroup
	//startGo := time.Now()
	//durationGo := time.Since(startGo)
	//for i := 0; i < 4; i++ {
	//	wg.Add(1)
	//	go inserter(i+20, &wg, db)
	//}
	//wg.Wait()
	//fmt.Println("with go - " + fmt.Sprintf("%.6f", durationGo))

	v2()

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

	sqlStatement := `
		INSERT INTO data (name, description)
		VALUES ($1, $2)`

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for j := 0; j < 4000; j++ {
		res, err := stmt.Exec(j, j)
		if err != nil || res == nil {
			log.Fatal(err)
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

func v2() {

	start := time.Now()
	duration := time.Since(start)

	sqlStatement := `
		INSERT INTO data (name, description)
		VALUES ($1, $2)`
	var (
		sStmt   string = sqlStatement
		gophers int    = 6
		entries int    = 40000
	)

	finishChan := make(chan int)

	for i := 0; i < gophers; i++ {
		go func(c chan int) {
			db, err := sql.Open("postgres", getDsn())
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			stmt, err := db.Prepare(sStmt)
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			for j := 0; j < entries; j++ {
				res, err := stmt.Exec(j, j)
				if err != nil || res == nil {
					log.Fatal(err)
				}
			}

			c <- 1
		}(finishChan)
	}

	finishedGophers := 0
	finishLoop := false
	for {
		if finishLoop {
			break
		}
		select {
		case n := <-finishChan:
			finishedGophers += n
			if finishedGophers == gophers {
				finishLoop = true
			}
		}
	}

	fmt.Println("with go - " + fmt.Sprintf("%.6f", duration))
}
