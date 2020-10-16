package main

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//	"time"
//)
//const (
//	host     = "postgresql"
//	port     = 5432
//	user     = "user"
//	password = "123"
//	dbname   = "db_main"
//)
//
//func getDsn() string {
//	return fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//}
//
//
//func main() {
//	fmt.Printf("StartTime: %v\n", time.Now())
//
//
//	sqlStatement := `
//		INSERT INTO data (name, description)
//		VALUES ($1, $2)`
//	var (
//		sStmt   string = sqlStatement
//		gophers int    = 5
//		entries int    = 10000
//	)
//
//	finishChan := make(chan int)
//
//	for i := 0; i < gophers; i++ {
//		go func(c chan int) {
//			db, err := sql.Open("postgres", getDsn())
//			if err != nil {
//				log.Fatal(err)
//			}
//			defer db.Close()
//
//			stmt, err := db.Prepare(sStmt)
//			if err != nil {
//				log.Fatal(err)
//			}
//			defer stmt.Close()
//
//			for j := 0; j < entries; j++ {
//				res, err := stmt.Exec(j, j)
//				if err != nil || res == nil {
//					log.Fatal(err)
//				}
//			}
//
//			c <- 1
//		}(finishChan)
//	}
//
//	finishedGophers := 0
//	finishLoop := false
//	for {
//		if finishLoop {
//			break
//		}
//		select {
//		case n := <-finishChan:
//			finishedGophers += n
//			if finishedGophers == 10 {
//				finishLoop = true
//			}
//		}
//	}
//
//	fmt.Printf("StopTime: %v\n", time.Now())
//}
