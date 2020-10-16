package main

import (
	"log"
	"time"
)

func v1() {
	start := time.Now()
	for i := 0; i < 4; i++ {
		inserterSimple(i + 20)
	}
	duration := time.Since(start)

	log.Print(duration.Seconds())
	//fmt.Println("without gorutunes - " . vv)
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

	for i := 0; i < 20000; i++ {
		_, err := db.Exec(sqlStatement, num, num)
		if err != nil {
			panic(err)
		}
	}
}
