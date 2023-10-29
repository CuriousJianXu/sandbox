package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"oproaster.com/sandbox/usecase"
)

func main() {
	connectionStr := "" // hidden for security reasons
	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Fatalln(err)
	}

	uc := usecase.New(db)
	uc.CrawlAndStoreTransactions()
}
