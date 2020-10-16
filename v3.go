package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func v3() {
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
