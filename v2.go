package main

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

func v2() {
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
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go inserter(i+20, &wg, db)
	}
	wg.Wait()
	durationGo := time.Since(startGo)

	log.Print(durationGo.Seconds())

	//fmt.Println("with gorutunes v1 - " + strconv.FormatFloat(durationGo.Seconds(), 'E', -1, 64))

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

	for j := 0; j < 20000; j++ {
		res, err := stmt.Exec(j, j)
		if err != nil || res == nil {
			log.Fatal(err)
		}
	}
}
