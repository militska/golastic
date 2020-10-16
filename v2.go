package main

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

const (
	gophers int = 6
	entries int = 40000
)

func loadData() {
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
	for i := 0; i < gophers; i++ {
		wg.Add(1)
		go insert(i+20, &wg, db)
	}
	wg.Wait()
	durationGo := time.Since(startGo)

	log.Print(durationGo.Seconds())
}

func insert(num int, group *sync.WaitGroup, db *sql.DB) {
	defer group.Done()

	sqlStatement := `
		INSERT INTO users (name, phone, company, birthday, address)
		VALUES ($1, $2, $3, $4, $5)`

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for j := 0; j < entries; j++ {
		user := generateUser()
		res, err := stmt.Exec(user.Name, user.Phone, user.Company, user.Birthday, user.Address)
		if err != nil || res == nil {
			log.Fatal(err)
		}
	}
}
